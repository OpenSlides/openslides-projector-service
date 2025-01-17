function setPageWidthVar() {
  const pageEl = document.getElementById(`projector-page`);
  const projectorWidth = +getComputedStyle(document.body).getPropertyValue(
    `--projector-width`,
  );
  const projectorAspectRatio =
    +getComputedStyle(document.body).getPropertyValue(
      `--projector-aspect-ratio-denominator`,
    ) /
    +getComputedStyle(document.body).getPropertyValue(
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

window.addEventListener("load", setPageWidthVar);
window.addEventListener("resize", setPageWidthVar);
