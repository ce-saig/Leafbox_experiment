$(function() {
    $(".dropable input[type=file]")
    .on("change", function(event) {
        readfiles(this.files, $(this).siblings('.file-preview'));
    })
    .on("drop", function(event) {
        event.preventDefault();  
        event.stopPropagation();
        var files = event.originalEvent.dataTransfer.files;
        this.files = files;
    });
    
});





//Dropable File Preview

var tests = {
      filereader: typeof FileReader != 'undefined',
      dnd: 'draggable' in document.createElement('span'),
      formdata: !!window.FormData,
      progress: "upload" in new XMLHttpRequest
    }, 
    support = {
      filereader: document.getElementById('filereader'),
      formdata: document.getElementById('formdata'),
      progress: document.getElementById('progress')
    },
    acceptedTypes = {
      'image/png': true,
      'image/jpeg': true,
      'image/gif': true
    };



function previewfile(file, element) {
    if (tests.filereader === true && acceptedTypes[file.type] === true) {
        var reader = new FileReader();
        reader.onload = function (event) {
            var image = new Image();
            image.src = event.target.result;
            image.style.width = '100%';
            //element.appendChild(image);
            $(element).html(image);
        };
      
        reader.readAsDataURL(file);
    }  else {
        //element.innerHTML += '<p>Uploaded ' + file.name + ' ' + (file.size ? (file.size/1024|0) + 'K' : '');
        $(element).html('<p>Uploaded ' + file.name + ' ' + (file.size ? (file.size/1024|0) + 'K' : ''));
        console.log(file);
    }
}

function readfiles(files, element) {
    var formData = tests.formdata ? new FormData() : null;
    for (var i = 0; i < files.length; i++) {
        if (tests.formdata) formData.append('file', files[i]);
        previewfile(files[i], element);
    }
}