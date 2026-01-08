export class ProjectorMotionBlock extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {
    this.observer = new ResizeObserver(() => {
      this.updateMotionNumberWidths();
      this.updateGridColumnCount();
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
      number.style.minWidth = maxNumberWidths[span.offsetLeft] + `px`;
    }
  }

  updateDisplayMotionTitle() {
    const motions = this.querySelectorAll(`.column-item`);
    const offsets = new Set();
    for (const motion of motions) {
      offsets.add(motion.offsetLeft);
    }

    const display = offsets.size > 1 ? `none` : null;
    const gridContainer = this.querySelector(`.motion-grid-container`);
    gridContainer.style.setProperty(`--title-display`, display);
  }

  updateGridColumnCount() {
    const gridContainer = this.querySelector(`.motion-grid-container`);
    const maxColumns = +gridContainer.style.getPropertyValue(`--max-columns`) || 3;
    for (let i = 0; i < maxColumns; i++) {
      gridContainer.style.setProperty(`--grid-column-count`, i + 1);

      if (this.offsetHeight >= gridContainer.offsetHeight) {
        return;
      }
    }
  }
}
