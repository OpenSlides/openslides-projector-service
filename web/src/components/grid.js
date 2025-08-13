export class OsGrid extends HTMLElement {
  constructor() {
    super();
    this.maxColumns = this.getAttribute(`max-columns`);
  }

  connectedCallback() {
    console.log(`test`);
  }
}

export class OsGridEntry extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {}

  isInViewport() {
    var rect = this.getBoundingClientRect();

    return (
      rect.top >= 0 &&
      rect.left >= 0 &&
      rect.bottom <= (window.innerHeight || document.documentElement.clientHeight) /* or $(window).height() */ &&
      rect.right <= (window.innerWidth || document.documentElement.clientWidth) /* or $(window).width() */
    );
  }
}
