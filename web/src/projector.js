import { EventSource } from 'eventsource';
import { setPageWidthVar } from './projector/scale.js';

/**
 * Creates a projector on the given element
 */
export function Projector(container, id, auth = () => ``) {
  const removeSizeListener = setPageWidthVar(container);
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
          Authorization: auth(),
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

  eventSource.addEventListener(`connected`, () => {
    console.debug(`connected`);
  });

  eventSource.addEventListener(`projector-replace`, (e) => {
    container.innerHTML = JSON.parse(e.data);
  });

  eventSource.addEventListener(`projection-updated`, (e) => {
    const data = JSON.parse(e.data);
    console.debug(`projection-updated`, data);

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
    removeSizeListener();
    eventSource.close();
  };
}
