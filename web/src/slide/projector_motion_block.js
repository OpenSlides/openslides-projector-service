export class ProjectorMotionBlock extends HTMLElement {
  MAX_COLUMNS = 3;

  constructor() {
    super();
  }

  connectedCallback() {
    this.observer = new ResizeObserver(() => {
      this.updateMotionNumberWidths();
      this.updateDisplayMotionTitle();
    });

    this.observer.observe(this);
  }

  disconnectedCallback() {
    this.observer.disconnect();
  }

  updateMotionNumberWidths() {
    const motionNumbers = this.querySelectorAll(`.motion-number`);
    let maxNumberWidths = {};
    for (const number of motionNumbers) {
      const span = number.querySelector(`span`);
      maxNumberWidths[span.offsetLeft] = Math.max(maxNumberWidths[span.offsetLeft] || 0, span.offsetWidth);
    }

    for (const number of motionNumbers) {
      const span = number.querySelector(`span`);
      number.style.width = maxNumberWidths[span.offsetLeft] + `px`;
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
}
