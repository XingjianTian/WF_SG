function renderClientList(data) {
    $.each(data, function (index, obj) {
        // render client status css tag style

        let clientStatusHtml = "";
        if (obj.enabled) {
            clientStatusHtml = `<br><a href = "#"><span class="badge badge-success" onclick="pauseClient('${obj.account}')">Working</span></a>`;
        }else{
            clientStatusHtml = `<br><a href = "#"><span  class="badge badge-danger" onclick="resumeClient('${obj.account}')">Disabled</span></a>`;
        }


        // render client allocated ip addresses
        let allocatedIpsHtml = "";
        $.each(obj.allocated_ips, function (index, ip) {
            allocatedIpsHtml += `<small class="badge badge-success">${ip.ip_address}</small>&nbsp;`;
        });

        // render client allowed ip addresses
        let allowedIpsHtml = "";
        $.each(obj.allowed_ips, function (index, ip) {
            allowedIpsHtml += `<small class="badge badge-success">${ip.ip_address}</small>&nbsp;`;
        });

        let html = `<div class="col-sm-4" id="client_${obj.account}">
                        <div class="small-box bg-gradient">
                        <div class="icon" disabled="true" readonly="true">
                            <i><img src="${obj.qrcode}"/></i>
                        </div>
                            <div class="inner">
                                    <h4 class ="fa fa-user-circle-o"><b>${obj.account}</b></h4>&nbsp;`
                               +clientStatusHtml+
            `<hr><p class="info-box-text"><b>IP Allocation</b></p>`
            + allocatedIpsHtml
            + `<hr>`
            + `<p class="info-box-text"><b>Allowed IPs</b></p>`
            + allowedIpsHtml
            + `<hr>`
            + `<div class="small-box-footer">
            <div class="btn-group">
                                <button type="button" class="btn btn-info btn-sm" data-toggle="modal"
                                        data-target="#modal_query_client" data-clientid="${obj.account}"
                                            data-clientname="${obj.account}">Details</button>
                                    <button type="button" class="btn btn-primary btn-sm" data-toggle="modal"
                                        data-target="#modal_edit_client" data-clientid="${obj.account}"
                                        data-clientname="${obj.account}">Edit</button>
                                    
                                    <button type="button" class="btn btn-danger btn-sm" data-toggle="modal"
                                        data-target="#modal_remove_client" data-clientid="${obj.account}"
                                        data-clientname="${obj.account}">Remove</button>
                                </div>
            </div>
                    </div>
                        </div>
                   </div>`;


        // add the client html elements to the list
        $('#client-list').append(html);


        let test2 = document.getElementById('client-list');
        let text = document.getElementsByName('client-list');
    });
}
