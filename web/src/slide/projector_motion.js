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
      if ([`changed`, `agreed`, `modified_final_version`].indexOf(this.mode) !== -1) {
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
    this.lineLength = this.getAttribute(`line-length`) ? +this.getAttribute(`line-length`) : null;
    this.firstLine = this.getAttribute(`first-line`) ? +this.getAttribute(`first-line`) : null;
    this.mode = this.getAttribute(`mode`);
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

  readChangeRecos() {
    const changes = [];
    this.querySelectorAll(`template.change-reco`).forEach(crEl => {
      changes.push({
        isTitleChange: false,
        identifier: crEl.getAttribute(`data-id`),
        lineFrom: +crEl.getAttribute(`data-line-from`),
        lineTo: +crEl.getAttribute(`data-line-to`),
        title: crEl.getAttribute(`data-change-title`),
        changeId: `r-${crEl.getAttribute(`data-id`)}`,
        changeType: crEl.getAttribute(`data-type`),
        changeNewText: crEl.getHTML().trim(),
        changeTitle: crEl.getAttribute(`data-change-title`)
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
            title: amendmentEl.getAttribute(`data-title`),
            identifier: amendmentEl.getAttribute(`data-number`),
            lineFrom: affectedLines.from,
            lineTo: affectedLines.to,
            changeId: `a-${amendmentEl.getAttribute(`data-id`)}-${pKey}`,
            changeType: `unknown`,
            changeNewText: affectedConsolidated,
            changeTitle: amendmentEl.getAttribute(`data-change-title`)
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
    const text = [];
    let lastLineTo = -1;
    for (let i = 0; i < changesToShow.length; i++) {
      if (changesToShow[i].lineTo > lastLineTo) {
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

    const changeHeader = [];
    if (HtmlDiff.changeHasCollissions(currentChange, changes)) {
      let style = `margin-left: 40px`;
      if (lineNumbering === `outside`) {
        style = `margin-right: 10px`;
      } else if (lineNumbering === `inside`) {
        style = `margin-left: 45px`;
      }

      changeHeader.push(
        `<span class="amendment-nr-n-icon"><span class="material-icons" style="${style}">warning</span>`
      );
    } else {
      let style = ` style="margin-left: 40px"`;
      if (lineNumbering === `outside`) {
        style = ``;
      } else if (lineNumbering === `inside`) {
        style = ` style="margin-left: 46px"`;
      }

      changeHeader.push(`<span class="amendment-nr-n-icon"${style}>`);
    }

    changeHeader.push(`<span class="amendment-nr">`);
    changeHeader.push(currentChange.changeTitle);
    changeHeader.push(`: </span></span>`);
    return changeHeader.join(``);
  }
}
