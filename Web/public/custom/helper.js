function renderClientList(data) {
    $.each(data, function (index, obj) {
        // render client status css tag style

        let clientStatusHtml = "";
        if (obj.enabled) {
            clientStatusHtml = `<a href = "#"><span class="badge badge-success" onclick="pauseClient('${obj.account}')">Working</span></a>&nbsp;`;
        }else{
            clientStatusHtml = `<a href = "#"><span  class="badge badge-danger" onclick="resumeClient('${obj.account}')">Disabled</span></a>&nbsp;`;
        }


        // render client allocated ip addresses
        let allocatedIpsHtml = "";
        $.each(obj.allocated_ips, function (index, ip) {
            allocatedIpsHtml += `<a href = "#"><span class="badge badge-success">${ip.ip_address}</span></a>&nbsp;`;
        });

        // render client allowed ip addresses
        let allowedIpsHtml = "";
        $.each(obj.allowed_ips, function (index, ip) {
            allowedIpsHtml += `<a href = "#"><span class="badge badge-success">${ip.ip_address}</span></a>&nbsp;`;
        });

        let html = `<div class="col-sm-4" id="client_${obj.account}">
                        <div class="info-box bg-gradient-white">
                        <div class="image">
                            <i><img src="${obj.qrcode}"/></i>
                        </div>
                            <div class="info-box-content">
                                    <span class ="fa fa-user-circle-o"><b>${obj.account}</b></span>`
            +clientStatusHtml+
            `<span class="info-box-text"><b>IP Allocation</b></span>`
            + allocatedIpsHtml

            + `<span class="info-box-text"><b>Allowed IPs</b></span>`
            + allowedIpsHtml

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
