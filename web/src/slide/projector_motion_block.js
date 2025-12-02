export class ProjectorMotionBlock extends HTMLElement {
  MAX_COLUMNS = 3;

  constructor() {
    super();
  }

  connectedCallback() {
    this.observer = new ResizeObserver(() => {
      this.updateMotionNumberWidths();
      this.updateGridColumnCount();
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
    const motionNumbers = this.querySelectorAll(`.motion-number`);
    const span = motionNumbers[0].querySelector(`span`);
    const maxNumberHeight = span.offsetHeight;

    const maxGridHeight = this.offsetHeight - 113; // the title is 113px high
    const numberOfMotionsPerColumn = (maxGridHeight / maxNumberHeight);

    const neededColumnAmount = Math.ceil(motionNumbers.length / numberOfMotionsPerColumn);
    
    const gridContainer = this.querySelector(`.grid-container`);
    if (neededColumnAmount > this.MAX_COLUMNS) {
      const extraColumnWidth = 100 / this.MAX_COLUMNS;
      const addtionalColumns = (neededColumnAmount * extraColumnWidth).toFixed(0)
      gridContainer.style.setProperty(`width`, `${addtionalColumns}%`);
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
