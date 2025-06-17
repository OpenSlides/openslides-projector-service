import { EventSource } from 'eventsource';
import { setPageWidthVar } from './projector/scale.js';
import { createProjectorClock } from './projector/clock.js';
import { ProjectorCountdown } from './slide/projector_countdown.js';
import { PdfViewer } from './components/pdf-viewer.js';

customElements.define("projector-countdown", ProjectorCountdown);
customElements.define("pdf-viewer", PdfViewer);

window.serverTime = () => new Date();

/**
 * Creates a projector on the given element
 */
export function Projector(host, id, auth = () => ``) {
  const container = host.attachShadow({ mode: `open` });
  const sizeListener = setPageWidthVar(container);
  const clock = createProjectorClock(container);
  let subscriptionUrl = `/system/projector/subscribe/${id}`;
  let needsInit = !container.childNodes.length;

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
          Authentication: auth(),
        },
      })
    },
  })
  eventSource.addEventListener(`settings`, (e) => {
    console.debug(`settings`, e.data);
  });

  eventSource.addEventListener(`deleted`, () => {
    console.debug(`deleted`);
  });

  eventSource.addEventListener(`connected`, (e) => {
    const timeOffset = +e.data - Math.floor(Date.now() / 1000);
    window.serverTime = () => {
      return new Date(Date.now() - (timeOffset * 1000));
    };
    clock.update();

    console.debug(`connected`);
  });

  eventSource.addEventListener(`projector-replace`, (e) => {
    container.innerHTML = JSON.parse(e.data);
    sizeListener.update();
    clock.update();
  });

  eventSource.addEventListener(`projection-updated`, (e) => {
    const data = JSON.parse(e.data);
    for (let id of Object.keys(data)) {
      let el = container.querySelector(`.slide[data-id="${id}"]`);
      if (!el) {
        el = container.querySelector(`#slides`).appendChild(document.createElement(`div`));
        el.classList.add(`slide`);
        el.dataset.id = id;
      }

      el.innerHTML = data[id];
    }
  });

  eventSource.addEventListener(`projection-deleted`, (e) => {
    console.debug(`projection-deleted`, e.data);

    container.querySelector(`.slide[data-id="${e.data}"]`)?.remove();
  });

  window.addEventListener(`unload`, () => {
    eventSource.close();
  });

  return () => {
    clearInterval(timeInterval);
    sizeListener.unregister();
    clock.unregister();
    eventSource.close();
  };
}
