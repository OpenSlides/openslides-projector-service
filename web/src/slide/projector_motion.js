import { HtmlDiff, LineNumbering } from '@openslides/motion-diff';

export class ProjectorMotionTitle extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {
    this.mode = this.getAttribute(`mode`);
    const crEl = this.querySelector(`template.change-reco`);

    const title = this.querySelector(`#content`).innerHTML;
    const container = document.createElement(`span`);
    if (crEl) {
      const changedTitle = crEl.getHTML().trim();
      if ([`changed`, `agreed`, `modified_final_version`].includes(this.mode)) {
        container.innerHTML = changedTitle;
      } else if (this.mode === `diff`) {
        container.innerHTML = HtmlDiff.diff(title, changedTitle);
      }
    } else {
      container.innerHTML = title;
    }

    this.appendChild(container);
  }
}

export class ProjectorMotionText extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {
    this.readAttributes();
    this.motionText = this.querySelector(`#content`).innerHTML;
    this.lineNumberedMotionText = null;

    this.changeRecos = this.readChangeRecos();
    this.amendmentChanges = this.readAmendmentChanges();

    switch (this.mode) {
      case `diff`:
        this.renderDiffView();
        break;
      case `changed`:
        this.renderChangeView();
        break;
      case `agreed`:
        this.renderFinalView();
        break;
      case `modified_final_version`:
        this.renderOriginalMotion();
        break;
      default:
        this.renderOriginalMotion();
    }
  }

  readAttributes() {
    this.lineLength = this.getAttribute(`line-length`) ? +this.getAttribute(`line-length`) : null;
    this.firstLine = this.getAttribute(`first-line`) ? +this.getAttribute(`first-line`) : null;
    this.mode = this.getAttribute(`mode`);
    this.i18n = this.getAttribute(`i18n`)
      ? JSON.parse(this.getAttribute(`i18n`))
      : {
          line: `Line`
        };
  }

  readChangeRecos() {
    const changes = [];
    this.querySelectorAll(`template.change-reco`).forEach(crEl => {
      changes.push({
        isTitleChange: false,
        identifier: crEl.getAttribute(`data-id`),
        lineFrom: +crEl.getAttribute(`data-line-from`),
        lineTo: +crEl.getAttribute(`data-line-to`),
        title: crEl.getAttribute(`data-change-title`) || ``,
        changeId: `r-${crEl.getAttribute(`data-id`)}`,
        changeType: crEl.getAttribute(`data-type`),
        changeNewText: crEl.getHTML().trim(),
        changeTitle: crEl.getAttribute(`data-change-title`) || ``
      });
    });

    return changes;
  }

  readAmendmentChanges() {
    const motionText = this.getLineNumberedMotionText();
    const motionParagraphs = LineNumbering.splitToParagraphs(motionText);

    const changes = [];
    this.querySelectorAll(`template.amendment`).forEach(amendmentEl => {
      const paragraphs = {};
      amendmentEl.content.querySelectorAll(`template.paragraph`).forEach(paragraphEl => {
        const original = motionParagraphs[+paragraphEl.getAttribute(`data-number`)];
        if (original === undefined) {
          return;
        }

        paragraphs[+paragraphEl.getAttribute(`data-number`)] = paragraphEl.getHTML().trim();
      });

      motionParagraphs.forEach((paragraph, pKey) => {
        const original = paragraph;

        let paragraphHasChanges = false;
        if (paragraphs[pKey] !== undefined) {
          // Add line numbers to newText, relative to the baseParagraph, by creating a diff
          // to the line numbered base version any applying it right away
          const diff = HtmlDiff.diff(paragraph, paragraphs[pKey]);
          paragraph = HtmlDiff.diffHtmlToFinalText(diff);
          paragraphHasChanges = true;
        }

        const affected = LineNumbering.getRange(paragraph);
        amendmentEl.content.querySelectorAll(`template.amendment-change-reco`).forEach(crEl => {
          const lineFrom = +crEl.getAttribute(`data-line-from`);
          const lineTo = +crEl.getAttribute(`data-line-to`);

          if (lineFrom >= affected.from && lineFrom <= affected.to) {
            paragraph = HtmlDiff.replaceLines(paragraph, crEl.getHTML().trim(), lineFrom, lineTo);

            // Reapply relative line numbers
            const diff = HtmlDiff.diff(motionParagraphs[pKey], paragraph);
            paragraph = HtmlDiff.diffHtmlToFinalText(diff);

            paragraphHasChanges = true;
          }
        });

        if (paragraphHasChanges) {
          const diff = HtmlDiff.diff(original, paragraph);
          const affectedLines = HtmlDiff.detectAffectedLineRange(diff);
          if (affectedLines === null) {
            return;
          }

          const affectedDiff = HtmlDiff.formatDiff(
            HtmlDiff.extractRangeByLineNumbers(diff, affectedLines.from, affectedLines.to)
          );
          const affectedConsolidated = HtmlDiff.diffHtmlToFinalText(affectedDiff);
          changes.push({
            isTitleChange: false,
            title: amendmentEl.getAttribute(`data-title`) || ``,
            identifier: amendmentEl.getAttribute(`data-number`),
            lineFrom: affectedLines.from,
            lineTo: affectedLines.to,
            changeId: `a-${amendmentEl.getAttribute(`data-id`)}-${pKey}`,
            changeType: `unknown`,
            changeNewText: affectedConsolidated,
            changeTitle: amendmentEl.getAttribute(`data-change-title`) || ``
          });
        }
      });
    });

    return changes;
  }

  getLineNumberedMotionText() {
    if (this.lineNumberedMotionText !== null) {
      return this.lineNumberedMotionText;
    }

    const config = {
      html: this.motionText
    };
    if (this.firstLine !== null) {
      config.firstLine = this.firstLine;
    }

    if (this.lineLength !== null) {
      config.lineLength = this.lineLength;
    }

    this.lineNumberedMotionText = LineNumbering.insert(config);
    return this.lineNumberedMotionText;
  }

  renderOriginalMotion() {
    const container = document.createElement(`div`);
    container.innerHTML = this.getLineNumberedMotionText();
    this.appendChild(container);
  }

  renderChangeView() {
    const container = document.createElement(`div`);
    container.innerHTML = HtmlDiff.getTextWithChanges(
      this.motionText,
      this.changeRecos,
      this.lineLength,
      false,
      null,
      this.firstLine
    );
    this.appendChild(container);
  }

  renderDiffView() {
    const motionText = this.getLineNumberedMotionText();
    const changesToShow = HtmlDiff.sortChangeRequests([...this.changeRecos, ...this.amendmentChanges]);
    changesToShow.sort(this.sortChangeRecBeforeAmend);
    const text = [];
    let lastLineTo = -1;
    for (let i = 0; i < changesToShow.length; i++) {
      if (changesToShow[i].lineFrom > lastLineTo + 1 && changesToShow[i].lineFrom > this.firstLine) {
        const changeFrom = changesToShow[i - 1] ? changesToShow[i - 1].lineTo + 1 : this.firstLine;
        text.push(
          HtmlDiff.extractMotionLineRange(
            motionText,
            {
              from: i === 0 ? this.firstLine : changeFrom,
              to: changesToShow[i].lineFrom - 1
            },
            true,
            this.lineLength
          )
        );
      }
      text.push(this.getChangeHeader(changesToShow, i));
      text.push(HtmlDiff.getChangeDiff(motionText, changesToShow[i], this.lineLength));

      lastLineTo = changesToShow[i].lineTo;
    }

    text.push(HtmlDiff.getTextRemainderAfterLastChange(motionText, changesToShow, this.lineLength));

    const container = document.createElement(`div`);
    container.innerHTML = text.join(``);
    this.appendChild(container);
  }

  renderFinalView() {
    const changesToShow = HtmlDiff.sortChangeRequests([...this.changeRecos, ...this.amendmentChanges]);
    changesToShow.sort(this.sortChangeRecBeforeAmend);

    const container = document.createElement(`div`);
    container.innerHTML = HtmlDiff.getTextWithChanges(
      this.motionText,
      changesToShow,
      this.lineLength,
      true,
      null,
      this.firstLine
    );
    this.appendChild(container);
  }

  getChangeHeader(changes, idx) {
    const lineNumbering = this.getAttribute(`line-numbering`);
    const currentChange = changes[idx];

    if (!HtmlDiff.changeHasCollissions(currentChange, changes) && currentChange.changeType != `unknown`) {
      return '';
    }

    const changeHeader = [];
    if (HtmlDiff.changeHasCollissions(currentChange, changes)) {
      let style = `margin-left: 40px`;
      if (lineNumbering === `outside`) {
        style = `margin-right: 15px`;
      } else if (lineNumbering === `inside`) {
        style = `margin-left: 45px`;
      }

      changeHeader.push(
        `<span class="amendment-nr-n-icon"><span class="material-icons" style="${style}">warning</span>`
      );
    } else {
      changeHeader.push(`<span class="amendment-nr-n-icon" style="margin-left: 40px">`);
    }

    if (currentChange.changeTitle) {
      changeHeader.push(`<span class="amendment-nr">`);
      changeHeader.push(currentChange.changeTitle);

      if (currentChange.changeType == `unknown`) {
        changeHeader.push(`: `);
      }
      changeHeader.push(`</span>`);
    }

    changeHeader.push(`</span>`);
    return changeHeader.join(``);
  }

  sortChangeRecBeforeAmend(a, b) {
    if (a.changeType == `unknown` && b.changeType == `unknown`) {
      return 1;
    } else if (a.changeType != `unknown` && b.lineFrom <= a.lineFrom && b.lineFrom >= a.lineTo && b != a) {
      return -1;
    }
    return 0;
  }
}

export class ProjectorMotionAmendment extends ProjectorMotionText {
  constructor() {
    super();
  }

  connectedCallback() {
    this.readAttributes();
    this.changeRecos = this.readChangeRecos();

    const leadMotionEl = this.querySelector(`template#lead-motion-text`);
    this.motionText = leadMotionEl.getHTML();

    this.lineNumberedMotionText = null;
    const motionText = this.getLineNumberedMotionText();
    this.motionParagraphs = LineNumbering.splitToParagraphs(motionText);

    this.changeParagraphs = {};
    this.querySelectorAll(`template.paragraph`).forEach(p => {
      this.changeParagraphs[p.getAttribute(`data-number`)] = p.getHTML();
    });

    this.render();
  }

  render() {
    const paragraphNumbers = Object.keys(this.changeParagraphs)
      .map(x => +x)
      .sort((a, b) => a - b);

    const amendmentParagraphs = paragraphNumbers
      .map(paraNo =>
        HtmlDiff.getAmendmentParagraphsLines(
          paraNo,
          this.motionParagraphs[paraNo],
          this.changeParagraphs[paraNo.toString()],
          this.lineLength,
          this.mode === `diff` ? this.changeRecos : undefined
        )
      )
      .filter(para => para !== null);

    const text = [];
    for (const p of amendmentParagraphs) {
      if (p.diffLineFrom === p.diffLineTo) {
        text.push(`<h3 class="amendment-line-header"><span>${this.i18n.line}</span> ${p.diffLineFrom}</h3>`);
      } else {
        text.push(
          `<h3 class="amendment-line-header"><span>${this.i18n.line}</span> ${p.diffLineFrom} - ${p.diffLineTo}</h3>`
        );
      }
      text.push(p.text);
    }

    const container = document.createElement(`div`);
    container.innerHTML = text.join(``);
    this.appendChild(container);
  }
}
