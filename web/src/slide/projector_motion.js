import { LineNumbering } from '@openslides/motion-diff';

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
    this.lineLength = this.getAttribute(`line-length`);
    this.firstLine = this.getAttribute(`first-line`);

    const config = {
      html: this.querySelector(`#content`).innerHTML
    };
    if (this.firstLine !== null) {
      config.firstLine = +this.firstLine;
    }

    if (this.lineLength !== null) {
      config.lineLength = +this.lineLength;
    }

    const container = document.createElement(`div`);
    container.innerHTML = LineNumbering.insert(config);
    this.appendChild(container);
  }
}
