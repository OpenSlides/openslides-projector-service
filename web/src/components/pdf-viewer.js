import { getDocument, GlobalWorkerOptions, PDFDocumentProxy } from "pdfjs-dist";
import * as PDFJSViewer from 'pdfjs-dist/web/pdf_viewer.mjs';

GlobalWorkerOptions.workerSrc = "/system/projector/static/lib/pdf.worker.mjs";

const CSS_UNITS = 96.0 / 72.0;

export class PdfViewer extends HTMLElement {
  constructor() {
    super();

    this.src = this.getAttribute(`src`);
  }

  connectedCallback() {
    this.pdfContainer = this.querySelector(`#pdf-container`);
    this.pdfContainer.style.position = `absolute`;
    const eventBus = new PDFJSViewer.EventBus();
    this.linkService = new PDFJSViewer.PDFLinkService({
      eventBus
    });
    this.findController = new PDFJSViewer.PDFFindController({
      eventBus,
      linkService: this.linkService
    });
    this.pdfViewer = new PDFJSViewer.PDFViewer({
      eventBus,
      container: this.pdfContainer,
      linkService: this.linkService,
      findController: this.findController
    });
    this.linkService.setViewer(this.pdfViewer);
    this.pdfViewer.currentScale = 2;
    
    eventBus.on("pagesinit", () => {
      // We can use pdfViewer now, e.g. let's change default scale.
      this.pdfViewer.currentScaleValue = "page-width";
    });

    getDocument({
      url: this.src,
      isEvalSupported: false,
      cMapUrl: '/system/projector/static/lib/cmaps/',
      cMapPacked: true
    }).promise.then((pdf) => {
      this.pdf = pdf;
      this.displayPdf(pdf);
    });
  }

  /**
   * @param {PDFDocumentProxy} pdf 
   */
  async displayPdf(pdf) {
    this.pdfViewer.setDocument(pdf);
    this.linkService.setDocument(pdf);
    this.findController.setDocument(pdf);

    await this.pdfViewer.firstPagePromise;
    this.updateSize();
  }

  async updateSize() {
    const page = await this.pdf.getPage(this.pdfViewer.currentPageNumber);

    const viewPort = page.getViewport({ scale: 1 });
    const scale = this.getScale(viewPort.width);

    this.pdfViewer.currentScale = scale;
  }

  getScale(viewportWidth) {
    const pdfContainerWidth = this.pdfContainer.clientWidth;

    if (
      pdfContainerWidth === 0 ||
      viewportWidth === 0
    ) {
      return 1;
    }

    return pdfContainerWidth / viewportWidth / CSS_UNITS;
  }
}
