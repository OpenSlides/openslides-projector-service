export class ProjectorCountdown extends HTMLElement {
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
    this.defaultTime = +this.getAttribute(`default-time`) || 0;
    this.countdownTime = +this.getAttribute(`countdown-time`) || 0;
    this.warningTime = +this.getAttribute(`warning-time`) || 0;
    this.running = this.getAttribute(`running`) === `true`;

    const displayType = this.getAttribute(`display-type`) || `onlyCountdown`;
    this.showTimeIndicator = displayType === `countdownAndTimeIndicator` || displayType === `onlyTimeIndicator`;
    this.showCountdown = displayType === `onlyCountdown` || displayType === `countdownAndTimeIndicator`;

    this.classList.add(`countdown-time-wrapper`);

    if (this.showTimeIndicator) {
      const timeIndicator = document.createElement(`div`);
      timeIndicator.id = `timeIndicator`;

      const timeIndicatorWrapper = document.createElement(`div`);
      timeIndicatorWrapper.classList.add(`time-indicator-wrapper`);
      timeIndicatorWrapper.appendChild(timeIndicator);
      this.appendChild(timeIndicatorWrapper);
    }

    if (this.showCountdown) {
      this.countdownEl = document.createElement(`div`);
      this.countdownEl.id = `countdown`;

      const countdownWrapper = document.createElement(`div`);
      countdownWrapper.classList.add(`countdown-wrapper`);
      countdownWrapper.appendChild(this.countdownEl);
      this.appendChild(countdownWrapper);
    }

    if (this.running) {
      this.updateCallback = setInterval(() => {
        this.updateComponent();
      }, 500);
    } else {
      this.updateComponent();
    }
  }

  updateComponent() {
    if (this.showCountdown) {
      this.countdownEl.innerText = this.countdownTimeFormatted;
    }

    const isNegative = this.seconds <= 0 && (this.defaultTime !== 0 || this.seconds < 0);
    const isWarning = this.defaultTime !== 0 && this.seconds <= this.warningTime && this.seconds > 0;
    if (isWarning) {
      this.classList.add('warning-time');
    }

    if (isNegative) {
      this.classList.add('negative-time');
    }
  }

  disconnectedCallback() {
    if (this.updateCallback) {
      clearInterval(this.updateCallback);
    }
  }
}
