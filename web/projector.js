import './projector/scale.js';

const eventSource = new EventSource(`/system/projector/subscribe/${window.location.href.substring(window.location.href.lastIndexOf('/') + 1)}`)
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
    let el = document.querySelector(`.slide[data-id="${id}"]`);
    if (!el) {
      el = document.getElementById(`slides`).appendChild(document.createElement(`div`));
      el.classList.add(`slide`);
      el.dataset.id = id;
    }

    el.innerHTML = data[id];
  }
});

eventSource.addEventListener(`projection-deleted`, (e) => {
  console.debug(`projection-deleted`, e.data);

  document.querySelector(`.slide[data-id="${e.data}"]`).remove();
});

window.addEventListener(`unload`, () => {
  eventSource.close();
});
