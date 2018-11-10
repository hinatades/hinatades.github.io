# ESLint
eslint js/src/
# js transform
babel --presets react,es2015 js/src/ -d js/build
# js package
browserify js/build/index.js -o bundle.js
# css package
cat css/*/* css/*.css | sed 's/..\/..\/images/images/g' > bundle.css
# done
echo; date; echo;
