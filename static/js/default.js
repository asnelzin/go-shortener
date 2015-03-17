(function(){

    function ajaxError(e) {
        if (!$('.form-group').hasClass('has-error'))
            $('.form-group').addClass('has-error');
        $('.control-label').text(e.responseJSON.error);
    }

    function ajaxSuccess(data) {
        formGroup = $('.form-group')
        if (formGroup.hasClass('has-error'))
            formGroup.removeClass('has-error');
        $('.control-label').text('');
        console.log(data);
        var url = location.protocol + '//' + location.host + '/' + data.shortUrl;
        $('#short-url').html('<a href="' + url + '">' + url + '</a>');
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