$(document).ready(function () {
    hljs.initHighlightingOnLoad();

    $('#sql tbody tr.description button').click(function() {
        $(this).find('i')
            .toggleClass('fas fa-eye')
            .toggleClass('fas fa-eye-slash');

        $(this).closest('#sql tbody tr').next().toggle();
    });

    $('#sql tbody button.close').click(function() {
        $(this).closest('#sql tbody tr').prev().find('button').click();
    });
});