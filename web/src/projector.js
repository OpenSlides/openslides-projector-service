import { EventSource } from 'eventsource';
import './projector/scale.js';

/**
 * Creates a projector on the given element
 */
export function Projector(container, id, auth = undefined) {
  const eventSource = new EventSource(`/system/projector/subscribe/${id}`, {
    fetch: (input, init) =>
      fetch(input, {
        ...init,
        headers: {
          ...init.headers,
          Authorization: auth,
        },
      }),
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

  eventSource.addEventListener(`projection-updated`, (e) => {
    const data = JSON.parse(e.data);
    console.debug(`projection-updated`, data);

    for (let id of Object.keys(data)) {
      let el = container.querySelector(`.slide[data-id="${id}"]`);
      if (!el) {
        el = container.getElementById(`slides`).appendChild(container.createElement(`div`));
        el.classList.add(`slide`);
        el.dataset.id = id;
      }

      el.innerHTML = data[id];
    }
  });

  eventSource.addEventListener(`projection-deleted`, (e) => {
    console.debug(`projection-deleted`, e.data);

    container.querySelector(`.slide[data-id="${e.data}"]`).remove();
  });

  window.addEventListener(`unload`, () => {
    eventSource.close();
  });
}
