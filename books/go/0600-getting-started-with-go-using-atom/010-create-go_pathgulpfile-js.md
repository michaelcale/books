Title: Create $GO_PATH/gulpfile.js
Id: 26864
Score: 1
Body:
    var gulp = require('gulp');
    var path = require('path');
    var shell = require('gulp-shell');
    
    var goPath = 'src/mypackage/**/*.go';
    
    
    gulp.task('compilepkg', function() {
      return gulp.src(goPath, {read: false})
        .pipe(shell(['go install <%= stripPath(file.path) %>'],
          {
              templateData: {
                stripPath: function(filePath) {
                  var subPath = filePath.substring(process.cwd().length + 5);
                  var pkg = subPath.substring(0, subPath.lastIndexOf(path.sep));
                  return pkg;
                }
              }
          })
        );
    });
    
    gulp.task('watch', function() {
      gulp.watch(goPath, ['compilepkg']);
    });
    
In the code above we defined a *compliepkg* task that will be triggered every time any go file in goPath (src/mypackage/) or subdirectories changes. the task will run the shell command go install changed_file.go

After creating the gulp file in go path and define the task open a command line and run:

> gulp watch

You'll se something like this everytime any file changes:
[![enter image description here][1]][1]


  [1]: https://i.stack.imgur.com/qwEvS.png
|======|
