import { EventSource } from 'eventsource';
import { setPageWidthVar } from './projector/scale.js';
import { createProjectorClock } from './projector/clock.js';
import { createOverlayOrganizer } from './projector/overlay.js';
import { OsIconContainer } from './components/icon-container.js';
import { ProjectorCountdown } from './slide/projector_countdown.js';
import { PdfViewer } from './components/pdf-viewer.js';
import { QrCode } from './components/qr-code.js';
import { ProjectorMotionBlock } from './slide/projector_motion_block.js';
import { ProjectorMotionAmendment, ProjectorMotionText, ProjectorMotionTitle } from './slide/projector_motion.js';
import { ProjectorPollChart } from './slide/poll_chart.js';

customElements.define('projector-countdown', ProjectorCountdown);
customElements.define('os-icon-container', OsIconContainer);
customElements.define('projector-motion-amendment', ProjectorMotionAmendment);
customElements.define('projector-motion-block', ProjectorMotionBlock);
customElements.define('projector-motion-title', ProjectorMotionTitle);
customElements.define('projector-motion-text', ProjectorMotionText);
customElements.define('projector-poll-chart', ProjectorPollChart);
customElements.define('pdf-viewer', PdfViewer);
customElements.define('qr-code', QrCode);

window.serverTime = () => new Date();

/**
 * Creates a projector on the given element
 */
export function Projector(host, id, auth = () => ``) {
  const container = host.attachShadow({ mode: `open` });
  const initContent = host.querySelector(`#current-content`).innerHTML;
  const sizeListener = setPageWidthVar(container);
  const clock = createProjectorClock(container);
  const overlayOrganizer = createOverlayOrganizer(container);
  let subscriptionUrl = `/system/projector/subscribe/${id}`;
  let needsInit = !initContent;
  if (!needsInit) {
    container.innerHTML = initContent;

    sizeListener.update();
    clock.update();
    overlayOrganizer.update();
  }

  const eventSource = new EventSource(subscriptionUrl, {
    fetch: (input, init) => {
      if (needsInit) {
        input.searchParams.set(`init`, `1`);
      }

      needsInit = true;
      return fetch(input, {
        ...init,
        headers: {
          ...init.headers,
          'ngsw-bypass': true,
          Authentication: auth()
        }
      });
    }
  });

  eventSource.addEventListener(`settings`, e => {
    const projectorContainer = container.querySelector(`#projector-container`);
    const settings = JSON.parse(e.data);
    const cssProperties = {
      '--projector-color': settings.Color,
      '--projector-background-color': settings.BackgroundColor,
      '--projector-header-background-color': settings.HeaderBackgroundColor,
      '--projector-header-font-color': settings.HeaderFontColor,
      '--projector-header-h1-color': settings.HeaderH1Color,
      '--projector-chyron-background-color': settings.ChyronBackgroundColor,
      '--projector-chyron-background-color2': settings.ChyronBackgroundColor2,
      '--projector-chyron-font-color': settings.ChyronFontColor,
      '--projector-chyron-font-color2': settings.ChyronFontColor2,
      '--projector-width': settings.Width,
      '--projector-aspect-ratio-numerator': settings.AspectRatioNumerator,
      '--projector-aspect-ratio-denominator': settings.AspectRatioDenominator,
      '--projector-scroll': settings.Scroll,
      '--projector-scale': settings.Scale
    };

    for (let prop in cssProperties) {
      if (cssProperties[prop] !== undefined) {
        projectorContainer.style.setProperty(prop, cssProperties[prop]);
      }
    }
  });

  eventSource.addEventListener(`deleted`, () => {
    console.debug(`deleted`);
  });

  eventSource.addEventListener(`connected`, e => {
    const timeOffset = +e.data - Math.floor(Date.now() / 1000);
    window.serverTime = () => {
      return new Date(Date.now() - timeOffset * 1000);
    };
    clock.update();

    console.debug(`connected`);
  });

  eventSource.addEventListener(`projector-replace`, e => {
    const html = JSON.parse(e.data);
    container.innerHTML = html;

    sizeListener.update();
    clock.update();
    overlayOrganizer.update();
  });

  eventSource.addEventListener(`projection-updated`, e => {
    const data = JSON.parse(e.data);

    for (let id of Object.keys(data)) {
      let el =
        container.querySelector(`#slides > [data-id="${id}"]`) ||
        container.querySelector(`.overlay-container > [data-id="${id}"]`);

      if (!el) {
        el = container.querySelector(`#slides`).appendChild(document.createElement(`div`));
        el.classList.add(`slide`);
        el.dataset.id = id;
      }

      el.innerHTML = data[id];
    }

    overlayOrganizer.update();
  });

  eventSource.addEventListener(`projection-deleted`, e => {
    console.debug(`projection-deleted`, e.data);

    container.querySelector(`#slides > [data-id="${e.data}"]`)?.remove();
    container.querySelector(`.overlay-container > [data-id="${e.data}"]`)?.remove();
  });

  window.addEventListener(`unload`, () => {
    eventSource.close();
  });

  return () => {
    sizeListener.unregister();
    clock.unregister();
    eventSource.close();
  };
}
