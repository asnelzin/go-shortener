(function(){



    var init = function(){
        $('form').on('submit', function(){
            $.ajax({
                url: form.attr('action'),
                data: form.serialize(),
                type: form.attr('method'),
                success: 
            })
            return false;
        });
    };
    
    $(init);
})()