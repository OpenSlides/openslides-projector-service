import * as esbuild from 'esbuild';

await esbuild.build({
  entryPoints: ['src/projector.js', 'src/projector.css', 'src/slide/*.css', 'src/slide/*.js'],
  bundle: true,
  minify: true,
  sourcemap: true,
  format: 'esm',
  target: ['chrome58', 'firefox57', 'safari11', 'edge18'],
  outdir: '../static/',
  external: ['*.woff']
});
