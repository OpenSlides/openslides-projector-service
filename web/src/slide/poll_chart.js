import { ArcElement, Chart, Colors, PieController } from 'chart.js';

Chart.register(PieController, ArcElement, Colors);

export class ProjectorPollChart extends HTMLElement {
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

    const dataset = { data };
    if (backgroundColor.length) {
      dataset.backgroundColor = backgroundColor;
    }

    const chart = new Chart(this.canvas, {
      type: 'doughnut',
      options: {
        hover: { mode: null },
        animations: false,
        maintainAspectRatio: false
      },
      data: {
        datasets: [dataset]
      }
    });

    for (let i = 0; i < chart.data.datasets[0].backgroundColor.length; i++) {
      const color = chart.data.datasets[0].backgroundColor[i];
      this.closest(`.result-wrapper`).style.setProperty(`--chart-bg-color-${i}`, color);
    }
  }
}
