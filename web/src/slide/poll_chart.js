import { ArcElement, Chart, Colors, PieController } from 'chart.js';

Chart.register(PieController, ArcElement, Colors);

export class ProjectorPollChart extends HTMLElement {
  config;

  connectedCallback() {
    this.config = JSON.parse(this.innerText.trim());
    this.innerHTML = ``;

    const shadow = this.attachShadow({ mode: 'open' });
    this.canvas = document.createElement('canvas');
    shadow.append(this.canvas);
    this.render();
  }

  render() {
    this.style.display = `block`;

    const data = [];
    const backgroundColor = [];
    for (const entry of this.config) {
      if (entry.color) {
        backgroundColor.push(window.getComputedStyle(this).getPropertyValue(entry.color));
      }
      data.push(entry.val);
    }

    new Chart(this.canvas, {
      type: 'doughnut',
      options: {
        hover: { mode: null }
      },
      data: {
        datasets: [
          {
            data,
            backgroundColor: backgroundColor.length ? backgroundColor : undefined
          }
        ]
      }
    });
  }
}
