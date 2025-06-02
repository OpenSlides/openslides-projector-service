export class ProjectorCountdown extends HTMLElement {
  defaultTime = 0;
  countdownTime = 0;
  running = false;
  updateCallback = null;

  get secondsRemaining() {
    const factor = this.defaultTime === 0 ? -1 : 1;
    if (this.running) {
      return Math.floor(this.countdownTime - window.serverTime() / 1000) * factor;
    }

    return this.countdownTime * factor;
  }

  /**
    * Updates the countdown time and string format it.
    */
  get countdownTimeFormatted() {
    this.seconds = this.secondsRemaining;

    const negative = this.seconds < 0;
    let seconds = this.seconds;
    if (negative) {
      seconds = -seconds;
    }

    const time = new Date(seconds * 1000);
    const m = Math.floor(+time / 1000 / 60).toString();
    const s = `0` + time.getSeconds();

    const timeString = (m.length < 2 ? `0` : ``) + m + `:` + s.slice(-2);
    if (negative) {
      return `-` + timeString;
    } else {
      return timeString;
    }
  }

  constructor() {
    super();
  }

  connectedCallback() {
    this.defaultTime = +this.getAttribute(`default-time`);
    this.countdownTime = +this.getAttribute(`countdown-time`);
    this.running = this.getAttribute(`running`) === `true`;

    this.updateCallback = setInterval(() => {
      this.innerText = this.countdownTimeFormatted;
    }, 500);
  }

  disconnectedCallback() {
    clearInterval(this.updateCallback);
  }
}
