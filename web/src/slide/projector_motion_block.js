export class ProjectorMotionBlock extends HTMLElement {
  MAX_COLUMNS = 3;

  constructor() {
    super();
  }

  connectedCallback() {
    this.observer = new ResizeObserver(() => {
      this.updateWidth(this.querySelectorAll(`.motion-number`));
      this.updateWidth(this.querySelectorAll(`.motion-detail`));
      this.updateGridColumnCount();
      this.updateDisplayMotionTitle();
    });

    this.observer.observe(this);
  }

  disconnectedCallback() {
    this.observer.disconnect();
  }

  updateWidth(nodeList) {
    let motionWidths = {};
    for (const number of nodeList) {
      const span = number.querySelector(`span`);
      motionWidths[span.offsetLeft] = Math.max(motionWidths[span.offsetLeft] || 0, span.offsetWidth);
    }

    for (const number of nodeList) {
      const span = number.querySelector(`span`);
      number.style.width = motionWidths[span.offsetLeft] + `px`;
    }
  }

  updateDisplayMotionTitle() {
    const motions = this.querySelectorAll(`.column-item`);
    const offsets = new Set();
    for (const motion of motions) {
      offsets.add(motion.offsetLeft);
    }

    const display = offsets.size > 1 ? `none` : null;
    for (const motion of motions) {
      motion.querySelector(`.motion-title`).style.display = display;
    }
  }

  updateGridColumnCount() {
    const gridContainer = this.querySelector(`.grid-container`);
    for (let i = 0; i < this.MAX_COLUMNS; i++) {
      gridContainer.style.setProperty(`--grid-column-count`, i + 1);

      if (this.offsetHeight >= gridContainer.offsetHeight) {
        return;
      }
    }
  }
}
