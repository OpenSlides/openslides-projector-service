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
customElements.define('projector-icon-container', OsIconContainer);
customElements.define('projector-motion-amendment', ProjectorMotionAmendment);
customElements.define('projector-motion-block', ProjectorMotionBlock);
customElements.define('projector-motion-title', ProjectorMotionTitle);
customElements.define('projector-motion-text', ProjectorMotionText);
customElements.define('projector-poll-chart', ProjectorPollChart);
customElements.define('projector-pdf-viewer', PdfViewer);
customElements.define('projector-qr-code', QrCode);

window.serverTime = () => new Date();

/**
 * Creates a projector on the given element
 */
export function Projector(host, id, auth = () => ``, config = {}) {
  config = Object.assign(
    {
      standalone: false,
      lang: null
    },
    config
  );

  const container = config.standalone ? host : host.attachShadow({ mode: `open` });
  const initContent = host.querySelector(`#current-content`)?.innerHTML;
  const sizeListener = setPageWidthVar(host, container);
  const clock = createProjectorClock(container);
  const overlayOrganizer = createOverlayOrganizer(container);
  let needsInit = !initContent;
  if (!needsInit) {
    container.innerHTML = initContent;

    sizeListener.update();
    clock.update();
    overlayOrganizer.update();
  }

  if (!id) {
    return () => {
      sizeListener.unregister();
      clock.unregister();
    };
  }

  const subscriptionUrl = `/system/projector/subscribe/${id}`;

  let eventSource = null;
  let reconnectDelay = 1000;
  const maxReconnectDelay = 30000;
  let reconnectTimer = null;
  let closedExplicitly = false;

  function createEventSource() {
    const es = new EventSource(subscriptionUrl, {
      fetch: (input, init) => {
        if (needsInit) {
          input.searchParams.set(`init`, `1`);
        }

        if (config.lang) {
          input.searchParams.set(`lang`, config.lang);
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

    attachListeners(es);
    return es;
  }

  function attachListeners(es) {
    es.addEventListener(`settings`, e => {
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

      console.debug(`[projector]`, `settings updated`, [e.data]);
    });

    es.addEventListener(`deleted`, () => {
      console.debug(`[projector]`, `deleted ${id}`);
    });

    es.addEventListener(`connected`, e => {
      const timeOffset = +e.data - Math.floor(Date.now() / 1000);
      window.serverTime = () => {
        return new Date(Date.now() - timeOffset * 1000);
      };
      clock.update();

      console.debug(`[projector]`, `connected ${id}`);
    });

    es.addEventListener(`projector-replace`, e => {
      console.debug(`[projector]`, `projector-replace`, [e.data]);

      const html = JSON.parse(e.data);
      container.innerHTML = html;

      sizeListener.update();
      clock.update();
      overlayOrganizer.update();
    });

    es.addEventListener(`projection-updated`, e => {
      console.debug(`[projector]`, `projection-updated`, [e.data]);

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

    es.addEventListener(`projection-deleted`, e => {
      console.debug(`[projector]`, `projection-deleted`, e.data);

      container.querySelector(`#slides > [data-id="${e.data}"]`)?.remove();
      container.querySelector(`.overlay-container > [data-id="${e.data}"]`)?.remove();
    });

    es.addEventListener(`error`, err => {
      console.warn(`[projector]`, `error for ${id}, trying to reconnect`);

      if (err.code >= 500 && err.code < 600) {
        scheduleReconnect();
      }
    });
  }

  async function probeHealthEndpoint() {
    try {
      const resp = await fetch(`/system/projector/health`, {
        method: 'GET',
        headers: {
          'ngsw-bypass': '1',
          Authentication: auth()
        },
        cache: 'no-store'
      });

      if (resp.ok) return true;

      return false;
    } catch (err) {
      console.warn(`[projector]`, `service unhealty`, err);
      return false;
    }
  }

  function scheduleReconnect() {
    if (closedExplicitly) return;
    if (reconnectTimer) return;

    reconnectTimer = setTimeout(async () => {
      reconnectTimer = null;

      const healthy = await probeHealthEndpoint();
      if (healthy) {
        reconnectDelay = 1000;
        eventSource = createEventSource();
      } else {
        console.warn(`[projector]`, `reconnecting in ${reconnectDelay}ms`);
        reconnectDelay = Math.min(reconnectDelay * 2, maxReconnectDelay);
        scheduleReconnect();
      }
    }, reconnectDelay);
  }

  eventSource = createEventSource();

  window.addEventListener(`unload`, () => {
    closedExplicitly = true;
    if (reconnectTimer) {
      clearTimeout(reconnectTimer);
      reconnectTimer = null;
    }
    eventSource?.close();
  });

  return () => {
    sizeListener.unregister();
    clock.unregister();
    closedExplicitly = true;
    if (reconnectTimer) {
      clearTimeout(reconnectTimer);
      reconnectTimer = null;
    }
    eventSource?.close();
  };
}
