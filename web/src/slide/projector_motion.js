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

  renderOriginalMotion() {
    const config = {
      html: this.motionText
    };
    if (this.firstLine !== null) {
      config.firstLine = this.firstLine;
    }

    if (this.lineLength !== null) {
      config.lineLength = this.lineLength;
    }

    const container = document.createElement(`div`);
    container.innerHTML = LineNumbering.insert(config);
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
}
