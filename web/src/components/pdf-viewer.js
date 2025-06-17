import { getDocument, GlobalWorkerOptions } from "pdfjs-dist";
import * as PDFJSViewer from 'pdfjs-dist/web/pdf_viewer.mjs';

GlobalWorkerOptions.workerSrc = "/system/projector/static/lib/pdf.worker.mjs";

export class PdfViewer extends HTMLElement {
  constructor() {
    super();

    this.src = this.getAttribute(`src`);
  }

  connectedCallback() {
    const pdfContainer = this.querySelector(`#pdf-container`);
    const eventBus = new PDFJSViewer.EventBus();
    const linkService = new PDFJSViewer.PDFLinkService({
      eventBus
    });
    const findController = new PDFJSViewer.PDFFindController({
      eventBus,
      linkService
    });
    const pdfViewer = new PDFJSViewer.PDFViewer({
      eventBus,
      container: pdfContainer,
      linkService: linkService,
      findController: findController
    });
    linkService.setViewer(pdfViewer);
    pdfViewer.currentScale = 2;

    getDocument({
      url: this.src,
      isEvalSupported: false,
      cMapUrl: '/system/projector/static/lib/cmaps/',
      cMapPacked: true
    }).promise.then((pdf) => {
      pdfViewer.setDocument(pdf);
      linkService.setDocument(pdf);
      findController.setDocument(pdf);
    });
    /*
    getDocument({ url: this.src }).promise.then((pdf) => {
      const context = pdfCanvas.getContext(`2d`);
      pdf.getPage(1).then(function(page) {
        const scale = 2;
        const viewport = page.getViewport({ scale: scale, });
        // Support HiDPI-screens.
        const outputScale = window.devicePixelRatio || 1;

        pdfCanvas.width = Math.floor(viewport.width * outputScale);
        pdfCanvas.height = Math.floor(viewport.height * outputScale);

        var transform = outputScale !== 1
          ? [outputScale, 0, 0, outputScale, 0, 0]
          : null;

        var renderContext = {
          canvasContext: context,
          transform: transform,
          viewport: viewport
        };
        page.render(renderContext).promise.then(() => {
          return page.getTextContent();
        }).then((textContent) => {
          console.log(textContent);
        });
        // you can now use *page* here
      });
    });
    */
  }
}
