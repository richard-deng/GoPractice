$(document).ready(function(){
    search_source();
    $.validator.addMethod("isMobile", function(value, element) {
        var length = value.length;
        var mobile = /^(1\d{10})$/;
        return this.optional(element) || (length == 11 && mobile.test(value));
    }, "请正确填写您的手机号码");

    $.validator.addMethod("isPhone", function(value, element) {
        var tel_pattern =  /^\d{3,4}-\d{7,8}(-\d{3,4})?$/;
        var mobile_pattern = /^(1\d{10})$/;
        return this.optional(element) || (tel_pattern.test(value)|| mobile_pattern.test(value));
    }, "请正确填写您的电话号码");

    $.validator.addMethod("isYuan", function(value, element) {
        var length = value.length;
        // var yuan  = /^([0-9]{1,6})\.([0-9]{1,2})$/;
        var yuan = /^([0-9]{1,8})(.([0-9]{1,2})){0,1}$/;
        return this.optional(element) || (length && yuan.test(value) && parseFloat(value) > 0);
    }, "请正确填写您的价格");

    $.validator.addMethod("isLessOne", function(value, element) {
        var length = value.length;
        var less_one  = /^(0)\.([0-9]{1,2})$/;
        return this.optional(element) || (length && less_one.test(value));
    }, "请正确填写您的比例");

    var table = $('#channelList').DataTable({
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
            var $channelList_length = $("#channelList_length");
            var $channelList_paginate = $("#channelList_paginate");
            var $page_top = $('.top');

            $page_top.addClass('row');
            $channelList_paginate.addClass('col-md-8');
            $channelList_length.addClass('col-md-4');
            $channelList_length.prependTo($page_top);
        },
        "ajax": function(data, callback, settings){
            var get_data = {
	           'page': Math.ceil(data.start / data.length) + 1,
	           'maxnum': data.length
            };

            var se_userid = window.localStorage.getItem('myid');
            get_data.se_userid = se_userid;

            var channel_name = $("#channelName").val();
            var phone_num = $("#s_phone_num").val();
            var is_prepayment = $("#s_is_prepayment").val();
            var is_valid = $("#s_is_valid").val();

            if(channel_name){
                get_data.channel_name = channel_name;
            }

            if(phone_num){
                get_data.phone_num = phone_num;
            }

            if(is_prepayment!=-1){
                get_data.is_prepayment = is_prepayment;
            }

            if(is_valid!=-1){
                get_data.is_valid = is_valid;
            }

            $.ajax({
	            url: '/api/channel/all',
	            type: 'POST',
	            dataType: 'json',
	            data: get_data,
	            success: function(data) {
                    var respcd = data.respcd;
                    if(respcd != '0000'){
                        $processing = $("#channelList_processing");
                        $processing.css('display', 'none');
                        var resperr = data.resperr;
                        var respmsg = data.respmsg;
                        var msg = resperr ? resperr : respmsg;
                        toastr.warning(msg);
                        return false;
                    }
	                detail_data = data.data;
	                num = detail_data.num;
	                callback({
	                    recordsTotal: num,
	                    recordsFiltered: num,
	                    data: detail_data.info
	                });
	            },
	            error: function(data) {
	                toastr.warning('获取数据异常');
	                return false;
	            }
            });
        },
        'columnDefs': [
            {
                targets: 7,
                data: '操作',
                render: function(data, type, full) {
                    var status = full.status;
                    var uid =full.userid;
                    var channel_id =full.id;
                    var msg = status ? '打开' : '关闭';
                    // var op = "<input type='button' class='btn btn-primary btn-sm setStatus' data-channelid="+uid+" value="+msg+ " data-status="+status+">";
                    var op = "<button type='button' class='btn btn-success btn-sm setStatus' data-channelid="+uid+" data-status="+status+">"+msg+"</button>";
                    // var view ="<input type='button' class='btn btn-info btn-sm viewEdit' data-uid="+uid+" value="+'查看'+ " data-channelid="+channel_id+">";
                    var view ="<button type='button' class='btn btn-warning btn-sm viewEdit' data-uid="+uid+" data-channelid="+channel_id+">"+'查看'+"</button>";
                    return op+view;
                }
            }
        ],
		'columns': [
				{ data: 'id' },
				{ data: 'channel_name' },
                //{ data: 'phone_num' },
				//{ data: 'contact_name' },
				{ data: 'training_amt_per' },
				{ data: 'divide_percent' },
				{ data: 'remain_times' },
				{ data: 'is_valid' },
				{ data: 'ctime' }
		],
        'oLanguage': {
            'sProcessing': '<span style="color:red;">加载中....</span>',
            'sLengthMenu': '每页显示_MENU_条记录',
            'sInfo': '显示 _START_到_END_ 的 _TOTAL_条数据',
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


	$("#channelCreate").click(function(){
        $("#channelCreateForm").resetForm();
        $("label.error").remove();
        get_all_rules();
        $('#create_divide_percent_div').show();
		$("#channelCreateModal").modal();
	});

    $("#channelNameSearch").click(function(){
        var channel_query_vt = $('#channel_query').validate({
           rules: {
               q_channel_name: {
                   required: false,
                   maxlength: 256
               },
               q_phone_num: {
                   required: false,
                   isMobile: '#s_phone_num'
               }
           },
           messages: {
               q_channel_name: {
                   required: '请输入渠道名称',
                   maxlength: $.validator.format("请输入一个长度最多是 {0} 的字符串")
               },
               q_phone_num: {
                   required: '请输入手机号'
               }
           },
           errorPlacement: function(error, element){
               var $error_element = element.parent().parent().next();
               $error_element.text('');
               error.appendTo($error_element);
           }
        });
        var ok = channel_query_vt.form();
        if(!ok){
            $("#query_label_error").show();
            $("#query_label_error").fadeOut(1400);
            return false;
        }
        $('#channelList').DataTable().draw();
    });

    $("#channelCreateSubmit").click(function(){

        var channel_create_vt = $('#channelCreateForm').validate({
            rules: {
                channel_name: {
                    required: true,
                    maxlength: 256
                },
                phone_num: {
                    required: true,
                    isMobile: '#phone_num'
                },
                address: {
                    required: true,
                    maxlength: 256
                },
                contact_name: {
                    required: true,
                    maxlength: 128
                },
                contact_phone: {
                    required: true,
                    isPhone: '#contact_phone'
                },
                contact_email: {
                    required: false,
                    email: true
                },
                training_amt_per: {
                    required: true,
                    isYuan: '#training_amt_per'
                },
                is_prepayment: {
                    required: true,
                    range:[0, 2]
                },
                /*
                email: {
                    required: false,
                    email: true
                },
                */
                divide_percent: {
                    required: true,
                    isLessOne: '#divide_percent'
                },
                create_rule: {
                    required: true,
                    minlength: 1
                }
            },
            messages: {
                channel_name: {
                    required: '请输入渠道名称',
                    maxlength: $.validator.format("请输入一个 长度最多是 {0} 的字符串")
                },
                phone_num: {
                    required: '请输入手机号'
                },
                address: {
                    required: '请输入地址',
                    maxlength: $.validator.format("请输入一个 长度最多是 {0} 的字符串")
                },
                contact_name: {
                    required: '请输入联系人姓名',
                    maxlength: $.validator.format("请输入一个 长度最多是 {0} 的字符串")
                },
                contact_phone: {
                    required: '请输入联系人手机号'
                },
                contact_email: {
                    email: "请输入正确格式的电子邮件"
                },
                training_amt_per: {
                    required: '请输入单次训练价格',
                    digits: "只能输入整数"
                },
                is_prepayment: {
                    required: '请选择结算模式',
                    range: $.validator.format("请输入一个介于 {0} 和 {1} 之间的值")
                },
                /*
                email: {
                    email: "请输入正确格式的电子邮件"
                },
                */
                divide_percent: {
                    required: '请正确填写比例'
                },
                create_rule: {
                    required: '请至少选择一个套餐',
                }
            },
            errorPlacement: function(error, element){
                if(element.is(':checkbox')){
                    error.appendTo(element.parent().parent().parent());
                } else {
                    error.insertAfter(element);
                }
            }
        });

        var ok = channel_create_vt.form();
        if(!ok){
            return false;
        }

        var post_data = {};
        var se_userid = window.localStorage.getItem('myid');
		var phone_num = $('#phone_num').val();
		var account_name = $('#account_name').val();
		var bank_name = $('#bank_name').val();
		var bank_account = $('#bank_account').val();
		var channel_name= $('#channel_name').val();
		var contact_name= $('#contact_name').val();
		var contact_phone= $('#contact_phone').val();
		var contact_email= $('#contact_email').val();
		var address= $('#address').val();
		var training_amt_per= $('#training_amt_per').val() * 100;
        var is_prepayment= $('.is_prepayment').val();
		var divide_percent= $('#divide_percent').val();

        //var email = $('#email').val();
        //var org_code = $('#org_code').val();
        //var license_id = $('#license_id').val();
        //var legal_person = $('#legal_person').val();
        //var business = $('#business').val();
        //var front_business = $('#front_business').val();

        post_data['se_userid'] = se_userid;
        post_data['login_name'] = phone_num;
		post_data['phone_num'] = phone_num;
		post_data['account_name'] = account_name;
		post_data['bank_name'] = bank_name;
		post_data['bank_account'] = bank_account;
		post_data['channel_name'] = channel_name;
		post_data['contact_name'] = contact_name;
		post_data['contact_phone'] = contact_phone;
		post_data['contact_email'] = contact_email;
		post_data['address'] = address;
		post_data['training_amt_per'] = training_amt_per;
		post_data['is_prepayment'] = is_prepayment;

		//post_data['business'] = business;
		//post_data['front_business'] = front_business;
        //post_data['email'] = email;
        //post_data['org_code'] = org_code;
        //post_data['license_id'] = license_id;
        //post_data['legal_person'] = legal_person;
		if(is_prepayment != 0){
		    if(!divide_percent){
		        toastr.warning('分成模式分成比例必填');
		        return false;
            }
            post_data['divide_percent'] = divide_percent;
        } else {
            post_data['rule'] = '';
            var rule = new Array();
            $('input[type=checkbox][name=create_rule]:checked').each(function(){
                console.log('xx val: '+$(this).val());
                rule.push($(this).val());
            });
            console.log(rule);
            post_data['rule'] = rule.join();
            console.log('selected checkbox arr: '+ post_data['rule']);
        }


        $.ajax({
	        url: '/channel_op/v1/api/channel_create',
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
                }
                else {
                    toastr.success('新建渠道成功');
                    search_source();
		            $("#channelCreateModal").modal('hide');
		            location.reload();
                    $('#channelList').DataTable().draw();
                }
	        },
	        error: function(data) {
                toastr.warning('请求异常');
	        }
        });
    });


    $(document).on('click', '.viewEdit', function(){
        $("label.error").remove();
        var uid = $(this).data('uid');
        var se_userid = window.localStorage.getItem('myid');
        var get_data = {
            'userid': uid,
            'se_userid': se_userid
        };
        $.ajax({
	        url: '/channel_op/v1/api/channel',
	        type: 'GET',
	        dataType: 'json',
	        data: get_data,
	        success: function(data) {
                var respcd = data.respcd;
                if(respcd != '0000'){
                    var resperr = data.resperr;
                    var respmsg = data.respmsg;
                    var msg = resperr ? resperr : respmsg;
                    toastr.warning(msg);
                }
                else {
                    var p_data = data.data.profile;
                    var ch_data = data.data.chn_data;
                    var u_data = data.data.u_data;
                    var crdata = data.data.crdata;
                    console.log('cr data :'+crdata);
                    $('#uid').text(uid);
                    $('#e_channel_id').val(ch_data.chnid);
                    $('#e_login_name').val(u_data.phone_num);
                    $('#e_phone_num').val(u_data.phone_num);
                    $('#e_channel_name').val(ch_data.channel_name);
                    //$('#e_create_time').val(ch_data.ctime);
                    //$('#e_legal_person').val(p_data.legal_person);
                    //$('#e_org_code').val(p_data.org_code);
                    //$('#e_license_id').val(p_data.license_id);
                    //$('#e_email').val(u_data.email);
                    //$('#e_business').val(p_data.business);
                    //$('#e_front_business').val(p_data.front_business);
                    $('#e_account_name').val(p_data.account_name);
                    $('#e_bank_account').val(p_data.bank_account);
                    $('#e_bank_name').val(p_data.bank_name);
                    $('#e_contact_name').val(p_data.contact_name);
                    $('#e_contact_phone').val(p_data.contact_phone);
                    $('#e_contact_email').val(p_data.contact_email);
                    $('#e_address').val(p_data.address);
                    $('#e_training_amt_per').val(ch_data.training_amt_per / 100);
                    $('#e_is_prepayment').val(ch_data.is_prepayment);
                    $('#e_divide_percent').val(ch_data.divide_percent);
                    // 添加套餐数据
                    edit_all_rules(crdata);
                    var is_prepayment = ch_data.is_prepayment;
                    if(is_prepayment==0){
                        $('#edit_divide_percent_div').hide();
                        $('#edit_training_amt_per').hide();
                        $('#edit_channel_rules').show();
                    }
                    else {
                        $('#edit_divide_percent_div').show();
                        $('#edit_training_amt_per').show();
                        $('#edit_channel_rules').hide();
                    }

                    $("#channelEditModal").modal();
                }
	        },
	        error: function(data) {
                toastr.warning('请求异常');
	        }
        });
    });


    $(document).on('click', '.setStatus', function(){
        var uid = $(this).data('channelid');
        var status = $(this).data('status');
        var value = status ? 0 : 1;
        var se_userid = window.localStorage.getItem('myid');
        var post_data = {
            'userid': uid,
            'state': value,
            'se_userid': se_userid
        };
        $.ajax({
	        url: '/channel_op/v1/api/channel_set_state',
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
                }
                else {
                    $('#channelList').DataTable().draw();
                    toastr.success('操作成功');
                }
	        },
	        error: function(data) {
                toastr.warning('请求异常');
	        }
        });
    });

    $('#channelEditSubmit').click(function(){

        var channel_edit_vt = $('#channelEditForm').validate({
            rules: {
                channel_name: {
                    required: true,
                    maxlength: 256
                },
                phone_num: {
                    required: true,
                    isMobile: '#phone_num'
                },
                address: {
                    required: true,
                    maxlength: 256
                },
                contact_name: {
                    required: true,
                    maxlength: 128
                },
                contact_phone: {
                    required: true,
                    isPhone: '#e_contact_phone'
                },
                contact_email: {
                    required: false,
                    email: true
                },
                training_amt_per: {
                    required: true,
                    isYuan: '#training_amt_per'
                },
                is_prepayment: {
                    required: true,
                    range:[0, 2]
                },
                email: {
                    required: false,
                    email: true
                },
                divide_percent: {
                    required: true,
                    isLessOne: '#e_divide_percent'
                },
                edit_rule: {
                    required: true,
                    minlength: 1
                }
            },
            messages: {
                channel_name: {
                    required: '请输入渠道名称',
                    maxlength: $.validator.format("请输入一个 长度最多是 {0} 的字符串")
                },
                phone_num: {
                    required: '请输入手机号'
                },
                address: {
                    required: '请输入地址',
                    maxlength: $.validator.format("请输入一个 长度最多是 {0} 的字符串")
                },
                contact_name: {
                    required: '请输入联系人姓名',
                    maxlength: $.validator.format("请输入一个 长度最多是 {0} 的字符串")
                },
                contact_phone: {
                    required: '请输入联系人手机号',
                },
                contact_email: {
                    email: "请输入正确格式的电子邮件"
                },
                training_amt_per: {
                    required: '请输入单次训练价格',
                    digits: "只能输入整数"
                },
                is_prepayment: {
                    required: '请选择结算模式',
                    range: $.validator.format("请输入一个介于 {0} 和 {1} 之间的值")
                },
                email: {
                    email: "请输入正确格式的电子邮件"
                },
                divide_percent: {
                    required: '请输入争取的比例'
                },
                edit_rule: {
                    required: '请至少选择一个套餐'
                }
            },
            errorPlacement: function(error, element){
                if(element.is(':checkbox')){
                    error.appendTo(element.parent().parent().parent());
                } else {
                    error.insertAfter(element);
                }
            }
        });
        var ok = channel_edit_vt.form();
        if(!ok){
            return false;
        }

	    var post_data = {};
        var uid = $('#uid').text();
        var se_userid = window.localStorage.getItem('myid');
        post_data['se_userid'] = se_userid;
        post_data['userid'] = uid;
		var phone_num = $('#e_phone_num').val();
		var email = $('#e_email').val();
		var org_code = $('#e_org_code').val();
		var license_id = $('#e_license_id').val();
		var legal_person = $('#e_legal_person').val();
		var account_name = $('#e_account_name').val();
		var bank_name = $('#e_bank_name').val();
		var bank_account = $('#e_bank_account').val();
		var channel_name= $('#e_channel_name').val();
		var contact_name= $('#e_contact_name').val();
		var contact_phone= $('#e_contact_phone').val();
		var contact_email= $('#e_contact_email').val();
		var address= $('#e_address').val();
		var training_amt_per= $('#e_training_amt_per').val() * 100;
		var divide_percent= $('#e_divide_percent').val();
		var is_prepayment= $('#e_is_prepayment').val();
		var business = $('#e_business').val();
		var front_business = $('#e_front_business').val();
        post_data['se_userid'] = se_userid;
		post_data['phone_num'] = phone_num;
		post_data['email'] = email;
		post_data['org_code'] = org_code;
		post_data['license_id'] = license_id;
		post_data['legal_person'] = legal_person;
		post_data['account_name'] = account_name;
		post_data['bank_name'] = bank_name;
		post_data['bank_account'] = bank_account;
		post_data['channel_name'] = channel_name;
		post_data['contact_name'] = contact_name;
		post_data['contact_phone'] = contact_phone;
		post_data['contact_email'] = contact_email;
		post_data['address'] = address;
		post_data['training_amt_per'] = training_amt_per;
		post_data['is_prepayment'] = is_prepayment;
		post_data['business'] = business;
		post_data['front_business'] = front_business;
		if(is_prepayment==1){
            // 分成模式
		    if(!divide_percent){
		        toastr.warning('分成模式分成比例必填');
		        return false;
            }
            post_data['divide_percent'] = divide_percent;
        } else {
            // 次卡模式
            post_data['rule'] = '';
            var rule = new Array();
            $('input[type=checkbox][name=edit_rule]:checked').each(function(){
                console.log('xx val: '+$(this).val());
                rule.push($(this).val());
            });
            console.log(rule);
            post_data['rule'] = rule.join();
            console.log('edit selected checkbox arr: '+ post_data['rule']);
        }

        $.ajax({
	        url: '/channel_op/v1/api/channel',
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
                }
                else {
                    toastr.success('保存修改成功');
                    $("#channelEditForm").resetForm();
                    $("#channelEditModal").modal('hide');
                    $('#channelList').DataTable().draw();
                }
	        },
	        error: function(data) {
                toastr.warning('请求异常');
	        }
        });
    });

    $('.is_prepayment').change(function(){
        var is_prepayment = $('.is_prepayment').val();
        if(is_prepayment === '0'){
            //次卡模式
            // $('#divide_percent').rules('remove');
            $('#divide_percent').next('label').remove();
            $('#create_divide_percent_div').hide();
            $('#create_training_amt_per').hide();
            $('#create_channel_rules').show();
        }else{
            //分成模式
            // $('#divide_percent').rules('add', { required: true, isLessOne: true, messages: {required: '请正确填写比例'}});
            $('#create_divide_percent_div').show();
            $('#create_training_amt_per').show();
            $('#create_channel_rules').hide();
        }
    });

    $('#e_is_prepayment').change(function(){
		var is_prepayment = $('#e_is_prepayment').val();
        if(is_prepayment === '0'){
            // 次卡模式
            // $('#e_divide_percent').rules('remove');
            $('#e_divide_percent').next('label').remove();
            $('#edit_divide_percent_div').hide();
            $('#edit_training_amt_per').hide();
            $('#edit_channel_rules').show();
        }else{
            // 分成模式
            // $('#e_divide_percent').rules('add', { required: true, isLessOne: true, messages: {required: '请正确填写比例'}});
            $('#edit_divide_percent_div').show();
            $('#edit_training_amt_per').show();
            $('#edit_channel_rules').hide();
        }
    });

});

function search_source() {
    var get_data = {};
    var se_userid = window.localStorage.getItem('myid');
    get_data['se_userid'] = se_userid;
    $.ajax({
        url: '/api/channel/names/all',
        type: 'GET',
        data: get_data,
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
                var subjects = new Array();
                console.log('subjects: ');
                console.log(data.data);
                for(var i=0; i<data.data.length; i++){
                    subjects.push(data.data[i])
                }
                $('#channelName').typeahead({source: subjects});
            }
        },
        error: function(data) {
            toastr.warning('请求异常');
        }
    });
}

function get_all_rules() {
    var get_data = {};
    var se_userid = window.localStorage.getItem('myid');
    get_data['se_userid'] = se_userid;

    var create_channel_rules =  $('#create_get_all_rules');
    create_channel_rules.html('');
    // $('#new_rules').html('');
    $.ajax({
        url: '/channel_op/v1/api/rules_list',
        type: 'GET',
        data: get_data,
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
                if(data.data.length == 0){
                    return false;
                }
                for(var i=0; i<data.data.length; i++){
                    var rule_id = data.data[i].id;
                    var rule_name = data.data[i].name;
                    var rule_description = data.data[i].description;
                    var option_str = $('<div class="checkbox"><label><input type="checkbox" name="create_rule" value="'+rule_id+'">'+rule_name + ' '+ rule_description+'</label></div>');
                    option_str.prependTo(create_channel_rules);
                }
            }
        },
        error: function(data) {
            msg = '请求异常';
            toastr.warning(msg);
        }
    });
    $('#create_channel_rules').hide();
}


function edit_all_rules(crdata) {
    var get_data = {};
    var se_userid = window.localStorage.getItem('myid');
    get_data['se_userid'] = se_userid;

    var edit_channel_rules =  $('#edit_get_all_rules');
    edit_channel_rules.html('');
    $.ajax({
        url: '/channel_op/v1/api/rules_list',
        type: 'GET',
        data: get_data,
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
                if(data.data.length == 0){
                    return false;
                }
                for(var i=0; i<data.data.length; i++){
                    var rule_id = data.data[i].id;
                    var rule_name = data.data[i].name;
                    var rule_description = data.data[i].description;
                    var option_str = $('<div class="checkbox"><label><input type="checkbox" name="edit_rule" value="'+rule_id+'">'+rule_name + ' '+ rule_description+'</label></div>');
                    option_str.prependTo(edit_channel_rules);
                }
                if(crdata.length > 0){
                    for(var i=0; i<crdata.length; i++){
                        op_str = 'input[type=checkbox][name="edit_rule"][value='+crdata[i]+']'
                        $(op_str).attr('checked', true);
                    }
                }

            }
        },
        error: function(data) {
            msg = '请求异常';
            toastr.warning(msg);
        }
    });
}
