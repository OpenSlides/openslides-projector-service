import QRCode from "qrcode";

export class QrCode extends HTMLElement {
  constructor() {
    super();

    this.text = this.getAttribute(`text`);
    this.size = this.getAttribute(`size`);
  }

  connectedCallback() {
    if (this.text) {
      const canvas = document.createElement(`canvas`);
      canvas.height = this.size;
      canvas.width = this.size;
      this.appendChild(canvas);

      QRCode.toCanvas(
        canvas,
        this.text,
        {
          width: this.size
        },
        (err) => { if (err) throw err; }
      );
    }
  }
}
