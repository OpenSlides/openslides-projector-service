export class FitColumns extends HTMLElement {
  constructor() {
    super();
  }


  connectedCallback() {
    this.style.display = `grid`;
    this.style.gap = `8px`;
    this.style.gridTemplateColumns = `repeat(1, 1fr)`;

    this.itemHeight = +this.getAttribute(`item-height`);
    this.maxColumns = +this.getAttribute(`max-columns`) || 3;
    this.maxHeight = +this.getAttribute(`max-height`);

    this.itemTemplate = this.querySelector(`template`);
    if (!this.itemTemplate) {
      console.warn(`FitColumns content template not found`);
      return;
    }

    this.initColumns();
  }

  /**
   * @returns HTMLElement
   */
  createColumn() {
    const column = document.createElement(`div`);
    column.classList.add(`column`);
    this.appendChild(column);
    this.columns.push(column);
    this.style.gridTemplateColumns = `repeat(${this.columns.length}, 1fr)`;

    return column;
  }

  initColumns() {
    this.maxHeight = +this.getAttribute(`max-height`) || this.offsetHeight;

    this.items = this.itemTemplate.content.querySelectorAll(`.column-item`);
    this.columns = [];
    this.createColumn();

    for (const item of this.items) {
      let currentColumn = this.columns[this.columns.length - 1];
      currentColumn.appendChild(item);

      if (currentColumn.offsetHeight > this.maxHeight && this.maxHeight > 0) {
        if (this.columns.length < this.maxColumns) {
          currentColumn = this.createColumn();
          currentColumn.appendChild(item);
        } else {
          let maxElements = currentColumn.childElementCount;
          for (let i = 0; i < this.columns.length - 1; i++) {
            if (this.columns[i].childElementCount < maxElements) {
              this.columns[i].appendChild(this.columns[i + 1].firstChild);
            }
          }
        }
      }
    }
  }

  updateColumns() {
    this.resetItems();
    this.initColumns();
  }

  resetItems() {
    for (const item of this.items) {
      this.itemTemplate.appendChild(item);
    }

    for (const column of this.querySelectorAll(`.column`)) {
      column.remove();
    }
    this.columns = [];
  }
}
