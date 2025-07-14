export function createProjectorClock(shadowDom) {
  let updateTimeout;
  function updateTime() {
    clearTimeout(updateTimeout);

    const time = window.serverTime();
    const clockEl = shadowDom.querySelector(`#clock-time`);
    if (clockEl) {
      const hour = `0${time.getHours()}`.substr(-2);
      const minute = `0${time.getMinutes()}`.substr(-2);
      clockEl.innerText = `${hour}:${minute}`;
    }

    const nextUpdateAt = new Date(Math.floor(time / 60000) * 60000 + 60000);
    setTimeout(updateTime, nextUpdateAt - time);
  }

  return {
    update() {
      updateTime();
    },
    unregister() {
      clearTimeout(updateTimeout);
    }
  };
}
