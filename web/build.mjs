import * as esbuild from 'esbuild';

let ctx = await esbuild.context({
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
    '.woff': 'file',
    '.woff2': 'file'
  },
  external: ['/assets/fonts/*.woff']
});

if (process.argv.indexOf(`--watch`) !== -1) {
  await ctx.watch();
  console.log('watching...');
} else {
  await ctx.rebuild();
  await ctx.dispose();
}
