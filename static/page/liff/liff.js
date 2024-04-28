function initializeLiff(myLiffId) {
    liff
        .init({
            liffId: myLiffId
        })
        .then(() => {
            initializeApp();
        })
        .catch((err) => {
            alert('啟動失敗。');
        });
}

function initializeApp() {
    var data = {
        accessToken: liff.getAccessToken()
    };
    fetch('https://thsr.onrender.com/v1/liff/verify', {
        body: JSON.stringify(data),
        headers: {
            'content-type': 'application/json'
        },
        method: 'POST'
    }).then(response => {
        return response.json();
    }).then(data => {
        var result = document.getElementById('result');
        result.innerHTML +=
            'scope:' + data.scope + '<br>' +
            'client_id:' + data.client_id + '<br>' +
            'expires_in:' + data.expires_in + '<br>';
        liff.getProfile().then(profile => {
            result.innerHTML +=
                'userID' + profile.userId
        })
    });
}

initializeLiff('2004698296-Vv4bvpZ7');
