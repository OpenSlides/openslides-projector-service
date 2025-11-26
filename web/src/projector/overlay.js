export function createOverlayOrganizer(shadowDom) {
  function organizeOverlays() {
    const container = shadowDom.querySelector('.overlay-container');
    if (!container) return;

    // Find all slides containing overlay elements
    shadowDom.querySelectorAll('.slide').forEach(slide => {
      const overlayElement = slide.querySelector('[data-is-overlay="true"]');
      if (overlayElement) {
        // Move the entire slide to countdown-container
        container.appendChild(slide);
      }
    });
  }

  return {
    update() {
      organizeOverlays();
    }
  };
}
