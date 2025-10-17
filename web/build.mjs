import { context } from 'esbuild';
import { copy } from 'esbuild-plugin-copy';

let ctx = await context({
  entryPoints: [
    'src/projector.js',
    'src/projector.css',
    'src/projector-page.css',
    'src/slide/*.css',
    'src/slide/*.js',
    'src/components/*.js'
  ],
  bundle: true,
  minify: true,
  sourcemap: true,
  format: 'esm',
  target: ['chrome58', 'firefox57', 'safari11', 'edge18'],
  outdir: '../static/',
  loader: {
    '.svg': 'file',
    '.gif': 'file',
    '.woff': 'file',
    '.woff2': 'file'
  },
  external: ['/assets/fonts/*.woff'],
  plugins: [
    copy({
      assets: {
        from: ['node_modules/pdfjs-dist/build/pdf.worker.mjs'],
        to: ['../static/lib']
      }
    }),
    copy({
      assets: {
        from: ['node_modules/pdfjs-dist/cmaps/*'],
        to: ['../static/lib/cmaps']
      }
    })
  ]
});

if (process.argv.indexOf(`--watch`) !== -1) {
  await ctx.watch();
  console.log('watching...');
} else {
  await ctx.rebuild();
  await ctx.dispose();
}
