export function setPageWidthVar(shadowDom) {
  const pageEl = shadowDom.host;
  function update() {
    const container = shadowDom.querySelector(`#projector-container`);
    console.log(pageEl, shadowDom, container);
    if (!container) {
      return;
    }

    const projectorWidth = +getComputedStyle(container).getPropertyValue(
      `--projector-width`,
    );
    const projectorAspectRatio =
      +getComputedStyle(container).getPropertyValue(
        `--projector-aspect-ratio-denominator`,
      ) /
      +getComputedStyle(container).getPropertyValue(
        `--projector-aspect-ratio-numerator`,
      );
    const projectorHeight = projectorWidth * projectorAspectRatio;
    const projectorPageAspectRatio = pageEl.offsetHeight / pageEl.offsetWidth;

    let containerWidth = pageEl.offsetWidth;
    if (projectorAspectRatio >= projectorPageAspectRatio) {
      containerWidth = pageEl.offsetHeight / projectorAspectRatio;
    }

    pageEl.style.setProperty("--projector-container-width", `${containerWidth}`);
    pageEl.style.setProperty(
      "--projector-container-height",
      `${(containerWidth / projectorWidth) * projectorHeight}`,
    );

    pageEl.style.setProperty("--projector-height", `${projectorHeight}`);
  }

  window.addEventListener("load", update);
  window.addEventListener("resize", update);

  return {
    update() {
      update();
    },
    unregister() {
      window.removeEventListener("load", update);
      window.removeEventListener("resize", update);
    }
  }
}
