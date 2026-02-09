export function setPageWidthVar(pageEl, shadowDom) {
  function update() {
    const container = shadowDom.querySelector(`#projector-container`);
    if (!container) {
      return;
    }

    const projectorWidth = +getComputedStyle(container).getPropertyValue(`--projector-width`);
    const projectorAspectRatio =
      +getComputedStyle(container).getPropertyValue(`--projector-aspect-ratio-denominator`) /
      +getComputedStyle(container).getPropertyValue(`--projector-aspect-ratio-numerator`);
    const projectorHeight = projectorWidth * projectorAspectRatio;
    const projectorPageAspectRatio = pageEl.offsetHeight / pageEl.offsetWidth;

    let containerWidth = pageEl.offsetWidth;
    if (projectorAspectRatio >= projectorPageAspectRatio) {
      containerWidth = pageEl.offsetHeight / projectorAspectRatio;
    }

    pageEl.style.setProperty('--projector-container-width', `${Math.floor(containerWidth)}`);
    pageEl.style.setProperty('--projector-container-height', `${(containerWidth / projectorWidth) * projectorHeight}`);

    pageEl.style.setProperty('--projector-height', `${projectorHeight}`);

    const headerHeight = shadowDom.querySelector(`#header`) ? 70 : 0;
    const footerHeight = shadowDom.querySelector(`#footer`) ? 35 : 0;
    const innerHeight = projectorHeight - headerHeight - footerHeight;
    pageEl.style.setProperty('--projector-inner-height', `${innerHeight}`);
  }

  window.addEventListener('load', update);
  window.addEventListener('resize', update);
  pageEl.addEventListener('resize', update);

  return {
    update() {
      update();
    },
    unregister() {
      window.removeEventListener('load', update);
      window.removeEventListener('resize', update);
      pageEl.removeEventListener('resize', update);
    }
  };
}
