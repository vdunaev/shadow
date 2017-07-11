function getWorkerStatusName(status) {
    switch(status) {
        case 0:
            return 'wait';
            break;
        case 1:
            return 'process';
            break;
        case 2:
            return 'busy';
            break;
        default:
            return 'unknown';
            break;
    }
}

function getTaskStatusName(status) {
    switch(status) {
        case 0:
            return 'wait';
            break;
        case 1:
            return 'process';
            break;
        case 2:
            return 'success';
            break;
        case 3:
            return 'fail';
            break;
        case 4:
            return 'fail by timeout';
            break;
        case 5:
            return 'kill';
            break;
        case 6:
            return 'wait repeat';
            break;
        default:
            return 'unknown';
            break;
    }
}

function update() {
    $.ajax({
        type: 'GET',
        url: '/workers/ajax/?action=stats',
        success: function(r){
            var tL = $('#listeners tbody').empty();
            var tW = $('#workers tbody').empty();
            var tT = $('#tasks tbody').empty();

            var listeners = r.listeners_count || 0;
            var workers = r.workers_count || 0;
            var tasks_wait = 0;

            if (Array.isArray(r.listeners)) {
                for (var i in r.listeners) {
                    var listener = r.listeners[i];

                    var lC = '<tr>'
                    + '<td>' + listener.name + '</td>'
                    + '<td>' + dateToString(listener.created) + '</td>'
                    + '<td>' + dateToString(listener.updated) + '</td>';
                    
                    if (listener.name != defaultListenerName) {
                        lC += '<td><div class="btn-group btn-group-sm">'
                            + '<button type="button" class="btn btn-danger btn-icon listener-remove" data-toggle="modal" data-target="#modal" data-modal-title="Confirm remove listener ' + listener.name + '" data-modal-callback="listenersRemove(\'' + listener.name + '\');"><i class="glyphicon glyphicon-trash" title="Remove"></i></button>'
                            + '</div></td>'
                    } else {
                        lC += '<td></td>';
                    }
                    
                    lC += '</tr>';
                    
                    tL.append(lC);
                }
                listeners = r.listeners.length
            }

            if (Array.isArray(r.workers)) {
                for (var i in r.workers) {
                    var worker = r.workers[i];
                    var task = worker.task;
                    var wC ='<tr>'
                        + '<td>' + worker.id + '</td>'
                        + '<td>' + dateToString(worker.created) + '</td>'
                        + '<td>' + getWorkerStatusName(worker.status) + '</td>';

                    if (task) {
                        wC += '<td><button type="button" class="btn btn-success btn-circle task-show" data-task="' + task.id + '"><i class="glyphicon glyphicon-eye-open"></i></button></td>';
                    } else {
                        wC += '<td></td>';
                    }

                    wC += '<td><div class="btn-group btn-group-sm">'
                        + '<button type="button" class="btn btn-info btn-icon worker-reset" data-toggle="modal" data-target="#modal" data-modal-title="Confirm reset worker #' + worker.id + '" data-modal-callback="workersReset(\'' + worker.id + '\');"><i class="glyphicon glyphicon-refresh" title="Reset"></i></button>'
                        + '<button type="button" class="btn btn-danger btn-icon worker-kill" data-toggle="modal" data-target="#modal" data-modal-title="Confirm kill worker #' + worker.id + '" data-modal-callback="workersKill(\'' + worker.id + '\');"><i class="glyphicon glyphicon-trash" title="Kill"></i></button>'
                        + '</div></td>'
                        + '</tr>';

                    tW.append(wC);

                    if (task) {
                        tW.append(
                            '<tr>'
                            + '<td colspan="5" style="display:none" id="task-' + i + '">'
                            + '<ul class="list-group">'
                            + '<li class="list-group-item"><span class="pull-right text-muted small"><em>' + task.id + '</em></span><strong>ID</strong><br /></li>'
                            + '<li class="list-group-item"><span class="pull-right text-muted small"><em>' + task.name + '</em></span><strong>Name</strong><br /></li>'
                            + '<li class="list-group-item"><span class="pull-right text-muted small"><em>' + getTaskStatusName(task.status) + '</em></span><strong>Status</strong><br /></li>'
                            + '<li class="list-group-item"><span class="pull-right text-muted small"><em>' + new Date(task.created).toLocaleString() + '</em></span><strong>Date</strong><br /></li>'
                            + '</ul>'
                            + '</td>'
                            + '</tr>'
                        )
                    }
                }
                workers = r.workers.length
            }

            if (Array.isArray(r.tasks_wait)) {
                for (var i in r.tasks_wait) {
                    var task = r.tasks_wait[i];
                    tT.append('<tr>'
                        + '<td>' + task.id + '</td>'
                        + '<td>' + task.name + '</td>'
                        + '<td>' + dateToString(task.created) + '</td>'
                        + '</tr>');
                }
                tasks_wait = r.tasks_wait.length
            }

            $('#listeners-count').text(listeners);
            $('#workers-count').text(workers);
            $('#tasks-title').text(tasks_wait + ' wait tasks');
            $('#tasks-wait-count').text(tasks_wait);

            $('#workers .task-show').click(function () {
                var e = $(this);
                var b = e.find('i');
                var i = e.data('task');
                
                if (b.hasClass('glyphicon-eye-open')) {
                    b.removeClass('glyphicon-eye-open');
                    b.addClass('glyphicon-eye-close');
                    $('#task-' + i).show('fast');
                } else {
                    b.removeClass('glyphicon-eye-close');
                    b.addClass('glyphicon-eye-open');
                    $('#task-' + i).hide('fast');
                }
            });
        }
    });
}

function workersKill(id) {
    $.ajax({
        type: 'POST',
        url: '/workers/ajax/?action=workers-kill',
        data: {
            id: id
        },
        success: update
    });
}

function workersReset(id) {
    $.ajax({
        type: 'POST',
        url: '/workers/ajax/?action=workers-reset',
        data: {
            id: id
        },
        success: update
    });
}

function listenersRemove(id) {
    $.ajax({
        type: 'POST',
        url: '/workers/ajax/?action=listeners-remove',
        data: {
            id: id
        },
        success: update
    });
}

$(function() {
    update();
    
    $('#workers-show').click(function () {
        $('#workers .task-show:has(i.glyphicon-eye-open)').click();
    });

    $('#workers-hide').click(function () {
        $('#workers .task-show:has(i.glyphicon-eye-close)').click();
    });

    $('#workers-add button[type=submit]').click(function () {
        $.ajax({
            type: 'POST',
            url: '/workers/ajax/?action=workers-add',
            data: {
                count: $('#workers-add-count').val()
            },
            success: update
        });
    });
    
    // 10 seconds
    setInterval(update, 1000 * 10);
});