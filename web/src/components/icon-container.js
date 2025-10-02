export class OsIconContainer extends HTMLElement {
  constructor() {
    super();
    this.icon = this.getAttribute(`icon`);
    this.iconClass = this.getAttribute(`iconClass`);
  }

  connectedCallback() {
    const iconElement = document.createElement(`span`);
    iconElement.classList.add(`material-icons`);
    if (this.iconClass) {
      iconElement.classList.add(this.iconClass);
    }
    iconElement.innerText = this.icon;
    this.prepend(iconElement);
  }
}
