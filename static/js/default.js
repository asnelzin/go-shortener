(function(){

    function ajaxError(e) {
        switch (e.status) {
            case 500:
               $(location).attr('href', '/500.html');
               break
            default:
                $('.form-group').toggleClass('has-error');
                $('.control-label').text(e.responseJSON.error);
        }
    }

    function ajaxSuccess(data) {
        $('.form-group').toggleClass('has-error');
        $('.control-label').text('');
        
        var url = location.protocol + '//' + location.host + '/s/' + data.shortUrl;
        $('#short-url').append('<a href="' + url + '">' + url + '</a>');
    }

    var init = function(){
        var form = $('form')
        form.on('submit', function(){
            $.ajax({
                url: form.attr('action'),
                data: form.serialize(),
                type: form.attr('method'),
                success: ajaxSuccess,
                error: ajaxError
            })
            return false;
        });

        $('input#url').keypress(function(e){
            if (e.which == 13) {
                form.submit();
                return false;
            }
        });
    };

    $(init);
})()