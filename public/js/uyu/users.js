/**
 * Created by admin on 2017/4/20.
 */
$(document).ready(function(){

    $.validator.addMethod("isMobile", function(value, element) {
        var length = value.length;
        var mobile = /^(1\d{10})$/;
        return this.optional(element) || (length == 11 && mobile.test(value));
    }, "请正确填写您的手机号码");


    $.validator.addMethod("isElevenNum", function(value, element) {
        var length = value.length;
        var mobile = /^([0-9]{11})$/;
        return this.optional(element) || (length == 11 && mobile.test(value));
    }, "请正确填写您的手机号码");

    $.validator.addMethod("PositiveNumber", function(value, element) {
        if(value <=0){
            return false;
        }
        else {
            return true;
        }
    }, "请正确填写您的次数");


    $('#userList').DataTable({
        "autoWidth": false,     //通常被禁用作为优化
        "processing": true,
        "serverSide": true,
        "paging": true,         //制指定它才能显示表格底部的分页按钮
        "info": true,
        "ordering": false,
        "searching": false,
        "lengthChange": true,
        "deferRender": true,
        "iDisplayLength": 10,
        "sPaginationType": "full_numbers",
        "lengthMenu": [[10, 40, 100],[10, 40, 100]],
        "dom": 'l<"top"p>rt',
        "fnInitComplete": function(){
            var $userList_length = $("#userList_length");
            var $userList_paginate = $("#userList_paginate");
            var $page_top = $('.top');

            $page_top.addClass('row');
            $userList_paginate.addClass('col-md-8');
            $userList_length.addClass('col-md-4');
            $userList_length.prependTo($page_top);
        },
        "ajax": function(data, callback, settings){
            var get_data = {
                'page': Math.ceil(data.start / data.length) + 1,
                'maxnum': data.length
            };

            var se_userid = window.localStorage.getItem('myid');
            get_data.se_userid = se_userid;
            var phone_num = $("#s_phone_num").val();

            if(phone_num){
                get_data.phone_num = phone_num;
            }

            var login_name = $("#s_login_name").val();
            if(login_name){
                get_data.login_name = login_name;
            }

            var nick_name = $("#s_nick_name").val();
            if(nick_name){
                get_data.nick_name = nick_name;
            }

            $.ajax({
                url: '/api/user/all',
                type: 'POST',
                dataType: 'json',
                data: get_data,
                success: function(data) {
                    var respcd = data.respcd;
                    if(respcd != '0000'){
                        $processing = $("#userList_processing");
                        $processing.css('display', 'none');
                        var resperr = data.resperr;
                        var respmsg = data.respmsg;
                        var msg = resperr ? resperr : respmsg;
                        toastr.warning(msg);
                        return false;
                    } else {
                        detail_data = data.data;
                        num = detail_data.num;
                        callback({
                            recordsTotal: num,
                            recordsFiltered: num,
                            data: detail_data.info
                        });
                    }
                },
                error: function(data) {
                    toastr.warning('请求数据异常');
                }

            });
        },
        'columnDefs': [
            {
                targets: 7,
                data: '操作',
                render: function(data, type, full) {
                    var userid = full.id;
                    var user_role = full.user_role;
                    var login_name = full.login_name;
                    var msg = '修改密码';
                    var op = "<button type='button' class='btn btn-success btn-sm modify-password' data-userid="+userid+">"+msg+"</button>";
                    if (user_role == 5 || user_role == 7){
                        var allocate_msg = "分配点数";
                        var allocate = "<button type='button' class='btn btn-primary btn-sm allocate-times' data-userid="+userid+ ' data-login_name='+login_name+">"+allocate_msg+"</button>";
                        return op + allocate;
                    } else {
                        return op;
                    }
                }
            }
        ],
        'columns': [
            { data: 'id' },
            { data: 'phone_num' },
            //{ data: 'username' },
            { data: 'login_name' },
            { data: 'nick_name' },
            { data: 'state' },
            { data: 'user_type' },
            //{ data: 'remain_times' },
            { data: 'ctime' }
        ],
        'oLanguage': {
            'sProcessing': '<span style="color:red;">加载中....</span>',
            'sLengthMenu': '每页显示_MENU_条记录',
            "sInfo": '显示 _START_到_END_ 的 _TOTAL_条数据',
            'sInfoEmpty': '没有匹配的数据',
            'sZeroRecords': '没有找到匹配的数据',
            'oPaginate': {
                'sFirst': '首页',
                'sPrevious': '前一页',
                'sNext': '后一页',
                'sLast': '尾页'
            }
        }
    });

    $("#userSearch").click(function(){

        var user_query_vt = $('#users_query').validate({
            rules: {
                q_phone_num: {
                    required: false,
                    isElevenNum: '#s_phone_num'
                },
                s_login_name: {
                    required: false,
                    maxlength: 255
                },
                s_nick_name: {
                    required: false,
                    maxlength: 128
                }
            },
            messages: {
                q_phone_num: {
                    required: '请输入手机号'
                },
                s_login_name: {
                    required: '请输入登录名称',
                    maxlength: $.validator.format("请输入一个长度最多是 {0} 的字符串")
                },
                s_nick_name: {
                    required: '请输入昵称',
                    maxlength: $.validator.format("请输入一个长度最多是 {0} 的字符串")
                }
            },
            errorPlacement: function(error, element){
                var $error_element = element.parent().parent().next();
                $error_element.text('');
                error.appendTo($error_element);
            }
        });
        var ok = user_query_vt.form();
        if(!ok){
            $("#query_label_error").show();
            $("#query_label_error").fadeOut(1400);
            return false;
        }
        $('#userList').DataTable().draw();
    });

    $(document).on('click', '.modify-password', function(){
        var userid = $(this).data('userid');
        $('#modify_userid').text(userid);
        $('#ModifyPassWordForm').resetForm();
        $('#ModifyPassWord').modal();
    });

    $(document).on('click', '.allocate-times', function () {
        $("#AllocateTimesForm").resetForm();
        $("label.error").remove();

        var userid = $(this).data('userid');
        var login_name = $(this).data('login_name');
        console.log('allocate times userid=' + userid +' login_name='+login_name);

        $('#user_login_name').val(login_name);
        $('#allocate_userid').text(userid);

        $('#allocate-times').modal();
    });

    $('.saveNewPassword').click(function () {
        var userid = $('#modify_userid').text();
        var new_password = $('#newPassword').val();
        var new_password_confirm = $('#newPasswordConfirm').val();

        if(new_password.length < 6 || new_password_confirm.length < 6){
            toastr.warning('密码长度至少六位');
            return false;
        }

        if(!new_password||!new_password_confirm||new_password!=new_password_confirm){
            toastr.warning('请检查密码是否一致');
            return false;
        }
        var se_userid = window.localStorage.getItem('myid');
        var post_data = {
            'se_userid': se_userid,
            'userid': userid,
            'password': md5(new_password),
            'echo_password': new_password,
        };
        $.ajax({
            url: '/api/user/password/change',
            type: 'POST',
            dataType: 'json',
            data: post_data,
            success: function(data) {
                var respcd = data.respcd;
                if(respcd != '0000'){
                    var resperr = data.resperr;
                    var respmsg = data.respmsg;
                    var msg = resperr ? resperr : respmsg;
                    toastr.warning(msg);
                    return false;
                } else {
                    $("#ModifyPassWord").modal('hide');
                    toastr.success('修改密码成功');
                }
            },
            error: function(data) {
                toastr.warning('请求数据异常');
            }
        });

    });

    $('.allocateTrainingTimes').click(function () {
        var post_url = '/channel_op/v1/api/platform_allocate_user';
        var allocate_vt = $('#AllocateTimesForm').validate({
            rules: {
                training_times: {
                    required: true,
                    digits: true,
                    PositiveNumber: '#training_times'
                }
            },
            messages: {
                training_times: {
                    required: "请输入训练次数",
                    digits: "只能输入整数"
                }
            }
        });
        var ok = allocate_vt.form();
        if(!ok){
            return false;
        }

        var userid = $('#allocate_userid').text();
        var training_times = $('#training_times').val();
        var post_data = {};
        var se_userid = window.localStorage.getItem('myid');
        post_data.se_userid = se_userid;
        post_data.userid = userid;
        post_data.training_times = training_times;

        $.ajax({
            url: post_url,
            type: 'POST',
            data: post_data,
            dataType: 'json',
            success: function(data) {
                var respcd = data.respcd;
                if(respcd != '0000'){
                    var resperr = data.resperr;
                    var respmsg = data.respmsg;
                    var msg = resperr ? resperr : respmsg;
                    toastr.warning(msg);
                }
                else {
                    toastr.success('分配训练点数成功');
                    $('#allocate-times').modal('hide');
                    $('#userList').DataTable().draw();
                }
            },
            error: function(data) {
                toastr.warning('请求异常');
            }
        });

    })
});
