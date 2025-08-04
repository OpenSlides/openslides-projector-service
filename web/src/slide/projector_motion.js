import { HtmlDiff, LineNumbering } from '@openslides/motion-diff';

export class ProjectorMotionTitle extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {
    const content = this.querySelector(`#content`).innerHTML;
    const container = document.createElement(`span`);
    container.innerHTML = content;

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
    this.changeRecos = this.readChangeRecos();

    switch (this.mode) {
      case `diff`:
        this.renderDiffView();
        break;
      case `changed`:
        this.renderChangeView();
        break;
      default:
        this.renderOriginalMotion();
    }
  }

  readChangeRecos() {
    const changeRecos = [];
    this.querySelectorAll(`template.change-reco`).forEach(crEl => {
      changeRecos.push({
        isTitleChange: false,
        identifier: crEl.getAttribute(`data-id`),
        lineFrom: +crEl.getAttribute(`data-line-from`),
        lineTo: +crEl.getAttribute(`data-line-to`),
        changeId: `r-${crEl.getAttribute(`data-id`)}`,
        changeType: crEl.getAttribute(`data-type`),
        changeNewText: crEl.getHTML().trim()
      });
    });

    return changeRecos;
  }

  getLineNumberedMotionText() {
    const config = {
      html: this.motionText
    };
    if (this.firstLine !== null) {
      config.firstLine = this.firstLine;
    }

    if (this.lineLength !== null) {
      config.lineLength = this.lineLength;
    }

    return LineNumbering.insert(config);
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
    const changesToShow = this.changeRecos;
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
              to: changesToShow[i].lineFrom - 1 || null
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

    text.push(
      HtmlDiff.getTextRemainderAfterLastChange(motionText, changesToShow, this.lineLength)
    );

    const container = document.createElement(`div`);
    container.innerHTML = text.join(``);
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
        `<span class="amendment-nr-n-icon"><mat-icon style="${style}">warning</mat-icon>`
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
    changeHeader.push(`TODO`);
    /*
    if (`amend_nr` in currentChange) {
      if (typeof currentChange.amend_nr === `string`) {
        changeHeader.push(currentChange.amend_nr);
      }
      if (currentChange.amend_nr === ``) {
        changeHeader.push(this.translate.instant(`Amendment`));
      }
    } else if (currentChange.getChangeType() === ViewUnifiedChangeType.TYPE_AMENDMENT) {
      const amendment = currentChange;
      changeHeader.push(amendment.getNumber(), ` - `, amendment.stateName);
    } else {
      if (currentChange.isRejected()) {
        changeHeader.push(this.translate.instant(`Change recommendation - rejected`));
      } else {
        changeHeader.push(this.translate.instant(`Change recommendation`));
      }
    }
    */
    changeHeader.push(`: </span></span>`);
    return changeHeader.join(``);
  }
}
