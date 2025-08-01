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
    const text = LineNumbering.insert({
      html: this.querySelector(`#content`).innerHTML
    });

    const container = document.createElement(`div`);
    container.innerHTML = text;

    this.appendChild(container);
  }
}
