export class ProjectorMotionBlock extends HTMLElement {
  MAX_COLUMNS = 3;

  constructor() {
    super();
  }

  connectedCallback() {
    this.observer = new ResizeObserver(() => {
      this.updateMotionNumberWidths();
      this.updateDisplayMotionTitle();
      this.updateHeight();
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

  updateHeight() {
    const gridContainer = this.querySelector(`.grid-container`);
    const motionNumbers = this.querySelectorAll(`.motion-number`);
    const maxNumberHeight = motionNumbers[0].querySelector(`span`).offsetHeight;
    const titleHeight = this.offsetParent.querySelector(`.slidetitle`).offsetHeight;

    const maxGridHeight = this.offsetHeight - titleHeight;
    const numberOfMotionsPerColumn = maxGridHeight / maxNumberHeight;

    const neededColumnAmount = Math.ceil(motionNumbers.length / numberOfMotionsPerColumn);
    const columnWidth = 100 / this.MAX_COLUMNS;
    const addtionalColumns = (neededColumnAmount * columnWidth).toFixed(0);

    gridContainer.style.setProperty(`--scroll-value`, `${(this.offsetWidth / 100) * columnWidth}px`);
    if (neededColumnAmount > this.MAX_COLUMNS) {
      gridContainer.style.setProperty(`width`, `${addtionalColumns}%`);
    }
  }
}
