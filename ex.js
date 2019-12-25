data.events.forEach(function (event) {

    var s, a, d, c, p, g, m = event.time;

    switch (t.ppks.forEach(function (n) {
        n && n.ppk_num == data.ppk_num && (g = n)
    }), event.data[0]) {
        case 8:
            a = "Идентификация пользователя", d = "Уведомление", c = "Пользователь №" + (event.data[1] - 15), s = "kz.svg", p = "identify", m -= 1, event.data[1] - 15 == 26 && (c = "Программирование ППК с клавиатуры", u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Программирование ППК с клавиатуры", +("826" + data.ppk_num))), event.data[1] - 15 == 27 && (c = "Программирование ППК через USB", u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Программирование ППК через USB", +("827" + data.ppk_num))), event.data[1] - 15 == 28 && (c = "Программирование ППК через Интернет", u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Программирование ППК через Интернет", +("828" + data.ppk_num))), event.data[1] - 15 == 29 && (c = "Управление с ПО Dunay control");
            break;
        case 9:
            a = "Снятие группы под принуждением", d = "Тревоги", s = "kz.svg", p = "alarm", m -= 1, u.alerts.createAlertsList("alarm", u.getName("group_name", data.ppk_num) + ": Снятие группы под принуждением", event.time, function () {
                u.nativeAudio.play("alarm").then(function () {
                    return u.nativeAudio.play("alarm")
                })
            }, function () {
                u.nativeAudio.stop("alarm")
            }), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Снятие группы под принуждением", +("9" + data.ppk_num), !0);
            break;
        case 88:
            a = "Норма шлейфа", d = "Уведомление", c = "Шлейф №" + (event.data[1] - 15), s = "norma.svg", p = "notification", m -= 9 - (event.data[1] - 15), e.plumes[event.data[1] - 15] = "Норма", u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Норма шлейфа: " + u.getName("Шлейф №" + (event.data[1] - 15), data.ppk_num), +("88" + (event.data[1] - 15) + data.ppk_num));
            break;
        case 105:
            a = "Восстановление сети 220В", d = "Уведомление", s = "220v_green.svg", p = "notification", e.state.v220 = !0, u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Восстановление сети 220В", +("105" + data.ppk_num));
            break;
        case 106:
            a = "Аккумулятор в норме", d = "Уведомление", s = "battery_green.svg", p = "notification", m += 1, e.state.battery = !0, u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Аккумулятор в норме", +("106" + data.ppk_num));
            break;
        case 108:
            a = "Открыта дверца ППК", d = "Тревоги", s = "door.svg", m += 2, p = "alarm", m += 1, e.state.door = !1, u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Открыта дверца ППК", +("108" + data.ppk_num), !0), u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": Открыта дверца ППК", event.time, function () {
                u.nativeAudio.play("alarm").then(function () {
                    return u.nativeAudio.play("alarm")
                })
            }, function () {
                u.nativeAudio.stop("alarm")
            });
            break;
        case 109:
            a = "Закрыта дверца ППК", d = "Уведомление", s = "door_green.svg", p = "notification", m += 2, e.state.door = !0, u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Закрыта дверца ППК", +("108" + data.ppk_num));
            break;
        case 63:
            if (d = "Уведомление", s = "kz.svg", p = "notification", m -= 20, 16 == event.data[1] && (a = "Рестарт ППК", u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Рестарт ППК", +("636" + data.ppk_num)), u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": Рестарт ППК", event.time)), 21 == event.data[1]) {
                console.log(JSON.parse(u.storage.get("settings")).ppks);
                if (!(v = JSON.parse(u.storage.get("settings")).ppks[r].ascribed)) {
                    a = "ППК приписан на ПЦН";
                    (f = JSON.parse(u.storage.get("settings"))).ppks[r].ascribed = !0, u.storage.set("settings", JSON.stringify(f)), console.log("ПРИПИСАН")
                }
            }
            if (22 == event.data[1]) {
                console.log(JSON.parse(u.storage.get("settings")).ppks);
                var v;
                if ((v = JSON.parse(u.storage.get("settings")).ppks[r].ascribed) || void 0 === v) {
                    a = "ППК отписан на ПЦН", console.log("ОТПИСАН");
                    var f;
                    (f = JSON.parse(u.storage.get("settings"))).ppks[r].ascribed = !1, u.storage.set("settings", JSON.stringify(f))
                }
            }
            break;
        case 64:
            a = "Взятие группы " + (event.data[1] - 15), d = "Управление", c = "group_name", s = "locked_red.svg", p = "lock", m -= 29 - (10 + event.data[1] - 15), e.groups[event.data[1] - 15] = !0, u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getGroupName(event.data[1] - 15, data.ppk_num) + " взята под охрану", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getGroupName(event.data[1] - 15, data.ppk_num) + " взята под охрану", +(event.data[1] - 15 + "" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("lock");
            break;
        case 72:
            a = "Снятие группы " + (event.data[1] - 15), d = "Управление", c = "group_name", s = "unlocked.svg", p = "unlock", u.alerts.clearErrorAlert(), m -= 29 - (10 + event.data[1] - 15), e.groups[event.data[1] - 15] = !1, u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getGroupName(event.data[1] - 15, data.ppk_num) + " снята с охраны", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getGroupName(event.data[1] - 15, data.ppk_num) + " снята с охраны", +(event.data[1] - 15 + "" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("unlock");
            break;
        case 240:
            a = "Взятие группы " + event.data[1] + ' в режиме "Остаюсь дома"', d = "Управление", c = "group_name", s = "locked_red.svg", p = "lock", u.alerts.clearErrorAlert(), m -= 29 - (10 + event.data[1]), e.groups[event.data[1]] = 2, u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getGroupName(event.data[1], data.ppk_num) + ' взята под охрану в режиме "Остаюсь дома"', event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getGroupName(event.data[1], data.ppk_num) + '  взята под охрану в режиме "Остаюсь дома"', +(event.data[1] + "" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("lock");
            break;
        case 80:
            a = "Обрыв шлейфа", d = "Тревоги", c = "Шлейф №" + (event.data[1] - 15), s = "kz.svg", p = "alarm", e.plumes[event.data[1] - 15] = "Обрыв", m -= 9 - (event.data[1] - 15 - 1), u.alerts.createAlertsList("alarm", "Обрыв шлейфа: " + u.getName("Шлейф №" + (event.data[1] - 15), data.ppk_num), event.time, function () {
                u.nativeAudio.play("alarm").then(function () {
                    return u.nativeAudio.play("alarm")
                })
            }, function () {
                u.nativeAudio.stop("alarm")
            }, u.getName("group_name", data.ppk_num)), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Обрыв шлейфа: " + u.getName("Шлейф №" + (event.data[1] - 15), data.ppk_num), +("80" + (event.data[1] - 15) + data.ppk_num), !0);
            break;
        case 104:
            a = "Отсутствие сети 220В", d = "Тревоги", s = "220v.svg", p = "alarm", m -= 1, e.state.v220 = !1, u.alerts.createAlertsList("alarm", u.getName("group_name", data.ppk_num) + ": Отсутствие сети 220В", event.time, function () {
                u.nativeAudio.play("alarm").then(function () {
                    return u.nativeAudio.play("alarm")
                })
            }, function () {
                u.nativeAudio.stop("alarm")
            }), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Отсутствие сети 220В", +("104" + data.ppk_num), !0);
            break;
        case 107:
            a = "Аккумулятор разряжен", d = "Тревоги", s = "battery.svg", p = "alarm", m += 1, e.state.battery = !1, u.alerts.createAlertsList("alarm", u.getName("group_name", data.ppk_num) + ": Аккумулятор разряжен", event.time, function () {
                u.nativeAudio.play("alarm").then(function () {
                    return u.nativeAudio.play("alarm")
                })
            }, function () {
                u.nativeAudio.stop("alarm")
            }), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Аккумулятор разряжен", +("107" + data.ppk_num), !0);
            break;
        case 112:
            a = "КЗ шлейфа", d = "Тревоги", c = "Шлейф №" + (event.data[1] - 15), s = "kz.svg", p = "alarm", e.plumes[event.data[1] - 15] = "КЗ", m -= 9 - (event.data[1] - 15 - 1), u.alerts.createAlertsList("alarm", "КЗ шлейфа: " + u.getName("Шлейф №" + (event.data[1] - 15), data.ppk_num), event.time, function () {
                u.nativeAudio.play("alarm").then(function () {
                    return u.nativeAudio.play("alarm")
                })
            }, function () {
                u.nativeAudio.stop("alarm")
            }, u.getName("group_name", data.ppk_num)), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": КЗ шлейфа: " + u.getName("Шлейф №" + (event.data[1] - 15), data.ppk_num), +("112" + (event.data[1] - 15) + data.ppk_num), !0);
            break;
        case 120:
            a = "Неисправность шлейфа", d = "Тревоги", c = "Шлейф №" + (event.data[1] - 15), s = "kz.svg", p = "alarm", e.plumes[event.data[1] - 15] = "Неисправность", m -= 9 - (event.data[1] - 15 - 1), u.alerts.createAlertsList("alarm", "Неисправность шлейфа: " + u.getName("Шлейф №" + (event.data[1] - 15), data.ppk_num), event.time, function () {
                u.nativeAudio.play("alarm").then(function () {
                    return u.nativeAudio.play("alarm")
                })
            }, function () {
                u.nativeAudio.stop("alarm")
            }, u.getName("group_name", data.ppk_num)), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Неисправность шлейфа: " + u.getName("Шлейф №" + (event.data[1] - 15), data.ppk_num), +("120" + (event.data[1] - 15) + data.ppk_num), !0);
            break;
        case 48:
            a = "Открыта крышка адаптера", d = "Уведомление", c = "Адаптер №" + (event.data[1] - 15), s = "kz.svg", p = "notification", e.adapters[event.data[1] - 15] = Object.assign(e.adapters[event.data[1] - 15] || {}, { door: !1 }), console.log("adapters:"), console.log(), u.alerts.createAlertsList("alarm", "Открыта крышка адаптера: Адаптер №" + (event.data[1] - 15), event.time, !0, !0, u.getName("group_name", data.ppk_num)), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Открыта крышка адаптера: Адаптер №" + (event.data[1] - 15), +("48" + (event.data[1] - 15) + data.ppk_num));
            break;
        case 50:
            a = "Закрыта крышка адаптера", d = "Уведомление", c = "Адаптер №" + (event.data[1] - 15), s = "kz.svg", p = "notification", e.adapters[event.data[1] - 15] = Object.assign(e.adapters[event.data[1] - 15] || {}, { door: !0 }), u.alerts.createAlertsList("alarm", "Закрыта крышка адаптера: Адаптер №" + (event.data[1] - 15), event.time, null, null, u.getName("group_name", data.ppk_num)), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Закрыта крышка адаптера: Адаптер №" + (event.data[1] - 15), +("50" + (event.data[1] - 15) + data.ppk_num));
            break;
        case 52:
            a = "Нет связи с адаптером", d = "Уведомление", c = "Адаптер №" + (event.data[1] - 15), s = "kz.svg", p = "notification", e.adapters[event.data[1] - 15] = Object.assign(e.adapters[event.data[1] - 15] || {}, { link: !1 }), u.alerts.createAlertsList("alarm", "Нет связи с адаптером: Адаптер №" + (event.data[1] - 15), event.time, !0, !0, u.getName("group_name", data.ppk_num)), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Нет связи с адаптером: Адаптер №" + (event.data[1] - 15), +("52" + (event.data[1] - 15) + data.ppk_num));
            break;
        case 54:
            a = "Восстановление связи с адаптером", d = "Уведомление", c = "Адаптер №" + (event.data[1] - 15), s = "kz.svg", p = "notification", e.adapters[event.data[1] - 15] = Object.assign(e.adapters[event.data[1] - 15] || {}, { link: !0 }), u.alerts.createAlertsList("alarm", "Восстановление связи с адаптером: Адаптер №" + (event.data[1] - 15), event.time, null, null, u.getName("group_name", data.ppk_num)), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Восстановление связи с адаптером: Адаптер №" + (event.data[1] - 15), +("54" + (event.data[1] - 15) + data.ppk_num));
            break;
        case 56:
            a = "Авария питания адаптера", d = "Уведомление", c = "Адаптер №" + (event.data[1] - 15), s = "kz.svg", p = "notification", e.adapters[event.data[1] - 15] = Object.assign(e.adapters[event.data[1] - 15] || {}, { battery: !1 }), u.alerts.createAlertsList("alarm", "Авария питания адаптера: Адаптер №" + (event.data[1] - 15), event.time, !0, !0, u.getName("group_name", data.ppk_num)), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Авария питания адаптера: Адаптер №" + (event.data[1] - 15), +("56" + (event.data[1] - 15) + data.ppk_num));
            break;
        case 58:
            a = "Питание в норме", d = "Уведомление", c = "Адаптер №" + (event.data[1] - 15), s = "kz.svg", p = "notification", e.adapters[event.data[1] - 15] = Object.assign(e.adapters[event.data[1] - 15] || {}, { battery: !0 }), u.alerts.createAlertsList("alarm", "Питание в норме: Адаптер №" + (event.data[1] - 15), event.time, null, null, u.getName("group_name", data.ppk_num)), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": Питание в норме: Адаптер №" + (event.data[1] - 15), +("50" + (event.data[1] - 15) + data.ppk_num));
            break;
        case 61:

            u.outputMessage && u.outputMessage();

            var h, _ = JSON.parse(u.storage.get("settings")).ppks;
            try {
                _.forEach(function (n) {
                    n && n.ppk_num == data.ppk_num && (h = n)
                })
            } catch (l) {
                console.error("e in fe", l)
            }

            switch (u.alerts.clearErrorAlert(), event.data[1]) {
                case 30:
                    "DunaySTK" != h.crd_type && (a = "Включение: ", c = "radio1", d = "Управление", s = "1", e.radio[1] = !0, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getRadioName("1", data.ppk_num) + " выключен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getRadioName("1", data.ppk_num) + " выключен", +("6130" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"));
                    break;
                case 31:
                    "DunaySTK" != h.crd_type && (a = "Выключение: ", c = "radio1", d = "Управление", s = "1", e.radio[1] = !1, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getRadioName("1", data.ppk_num) + " выключен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getRadioName("1", data.ppk_num) + " выключен", +("6131" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"));
                    break;
                case 32:
                    "DunaySTK" != h.crd_type && (a = "Включение: ", c = "radio2", d = "Управление", s = "2", e.radio[2] = !0, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getRadioName("2", data.ppk_num) + " выключен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getRadioName("2", data.ppk_num) + " выключен", +("6132" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"));
                    break;
                case 33:
                    "DunaySTK" != h.crd_type && (a = "Выключение: ", c = "radio2", d = "Управление", s = "2", e.radio[2] = !1, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getRadioName("2", data.ppk_num) + " выключен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getRadioName("2", data.ppk_num) + " выключен", +("6133" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"));
                    break;
                case 34:
                    "DunaySTK" != h.crd_type && (a = "Включение: ", c = "radio3", d = "Управление", s = "3", e.radio[3] = !0, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getRadioName("3", data.ppk_num) + " выключен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getRadioName("3", data.ppk_num) + " выключен", +("6134" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"));
                    break;
                case 35:
                    "DunaySTK" != h.crd_type && (a = "Выключение: ", c = "radio3", d = "Управление", s = "3", e.radio[3] = !1, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getRadioName("3", data.ppk_num) + " выключен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getRadioName("3", data.ppk_num) + " выключен", +("6135" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"));
                    break;
                case 17:
                    "DunaySTK" != h.crd_type && (a = "Выключение: ", c = "0", d = "Управление", s = "0", e.outputs[0] = !1, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getName("0", data.ppk_num) + " выключен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getName("0", data.ppk_num) + " выключен", +("6117" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"), g.outputs[0].pulse && u.pulse && (n = !1));
                    break;
                case 16:
                    "DunaySTK" != h.crd_type && (a = "Включение: ", c = "0", d = "Управление", s = "0", e.outputs[0] = !0, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getName("0", data.ppk_num) + " включен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getName("0", data.ppk_num) + " включен", +("6116" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"), g.outputs[0].pulse && u.pulse && (n = !1));
                    break;
                case 11:
                    "DunaySTK" != h.crd_type && "Dunay8l" != h.crd_type && (a = "Выключение: ", c = "1", d = "Управление", s = "1", e.outputs[1] = !1, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getName("1", data.ppk_num) + " выключен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getName("1", data.ppk_num) + " выключен", +("6111" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"), g.outputs[1].pulse && u.pulse && (n = !1));
                    break;
                case 10:
                    "DunaySTK" != h.crd_type && (a = "Включение: ", c = "1", d = "Управление", s = "1", e.outputs[1] = !0, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getName("1", data.ppk_num) + " включен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getName("1", data.ppk_num) + " включен", +("6110" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"), g.outputs[1].pulse && u.pulse && (n = !1));
                    break;
                case 13:
                    "DunaySTK" != h.crd_type && (a = "Выключение: ", c = "2", d = "Управление", s = "2", e.outputs[2] = !1, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getName("2", data.ppk_num) + " выключен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getName("2", data.ppk_num) + " выключен", +("6113" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"), g.outputs[2].pulse && u.pulse && (n = !1));
                    break;
                case 12:
                    "DunaySTK" != h.crd_type && (a = "Включение: ", c = "2", d = "Управление", s = "2", e.outputs[2] = !0, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getName("2", data.ppk_num) + " включен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getName("2", data.ppk_num) + " включен", +("6112" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"), g.outputs[2].pulse && u.pulse && (n = !1));
                    break;
                case 15:
                    "DunaySTK" != h.crd_type && "Dunay8l" != h.crd_type && (a = "Выключение: ", c = "3", d = "Управление", s = "3", e.outputs[3] = !1, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getName("3", data.ppk_num) + " выключен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getName("3", data.ppk_num) + " выключен", +("6115" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"), g.outputs[3].pulse && u.pulse && (n = !1));
                    break;
                case 14:
                    "DunaySTK" != h.crd_type && "Dunay8l" != h.crd_type && (a = "Включение: ", c = "3", d = "Управление", s = "3", e.outputs[3] = !0, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getName("3", data.ppk_num) + " включен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getName("3", data.ppk_num) + " включен", +("6114" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"), g.outputs[3].pulse && u.pulse && (n = !1));
                    break;
                case 6:
                    a = "Выключение: ", c = "relay2", d = "Управление", s = "relay2", e.outputs.relay2 = !1, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getName("relay2", data.ppk_num) + " выключен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getName("relay2", data.ppk_num) + " выключен", +("616" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"), g.outputs.relay2.pulse && u.pulse && (n = !1);
                    break;
                case 5:
                    a = "Включение: ", c = "relay2", d = "Управление", s = "relay2", e.outputs.relay2 = !0, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getName("relay2", data.ppk_num) + " включен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getName("relay2", data.ppk_num) + " включен", +("615" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"), g.outputs.relay2.pulse && u.pulse && (n = !1);
                    break;
                case 4:
                    a = "Выключение: ", c = "uk3", d = "Управление", s = "uk3", e.outputs.uk3 = !1, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getName("uk3", data.ppk_num) + " выключен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getName("uk3", data.ppk_num) + " выключен", +("614" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"), g.outputs.uk3.pulse && u.pulse && (n = !1);
                    break;
                case 3:
                    a = "Включение: ", c = "uk3", d = "Управление", s = "uk3", e.outputs.uk3 = !0, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getName("uk3", data.ppk_num) + " включен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getName("uk3", data.ppk_num) + " включен", +("613" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"), g.outputs.uk3.pulse && u.pulse && (n = !1);
                    break;
                case 2:
                    a = "Выключение: ", c = "uk2", d = "Управление", s = "uk2", e.outputs.uk2 = !1, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getName("uk2", data.ppk_num) + " выключен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getName("uk2", data.ppk_num) + " выключен", +("612" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"), g.outputs.uk2.pulse && u.pulse && (n = !1);
                    break;
                case 1:
                    a = "Включение: ", c = "uk2", d = "Управление", s = "uk2", e.outputs.uk2 = !0, p = "control", u.showAlert && (o != data.ppk_num && !cordova.plugins.backgroundMode.isActive() || event.time < u.connection_date) && u.alerts.createAlertsList("alert", u.getName("group_name", data.ppk_num) + ": " + u.getName("uk2", data.ppk_num) + " включен", event.time), u.alerts.createNotification(u.getName("group_name", data.ppk_num) + ": " + u.getName("uk2", data.ppk_num) + " включен", +("611" + data.ppk_num)), event.time > u.connection_date && u.alerts.response("output"), g.outputs.uk2.pulse && u.pulse && (n = !1)
            }
            break;
        case 5:
            a = "Открыта крышка", d = "Уведомление", c = "Беспроводной датчик №" + (event.data[1] + 1), s = "kz.svg", p = "notification", e.sensors[event.data[1] + 1] = Object.assign(e.sensors[event.data[1] + 1] || {}, { door: !1 });
            break;
        case 4:
            a = "Закрыта крышка", d = "Уведомление", c = "Беспроводной датчик №" + (event.data[1] + 1), s = "kz.svg", p = "notification", e.sensors[event.data[1] + 1] = Object.assign(e.sensors[event.data[1] + 1] || {}, { door: !0 });
            break;
        case 6:
            a = "Нет связи", d = "Уведомление", c = "Беспроводной датчик №" + (event.data[1] + 1), s = "kz.svg", p = "notification", e.sensors[event.data[1] + 1] = Object.assign(e.sensors[event.data[1] + 1] || {}, { link: !1 });
            break;
        case 7:
            a = "Восстановление связи", d = "Уведомление", c = "Беспроводной датчик №" + (event.data[1] + 1), s = "kz.svg", p = "notification", e.sensors[event.data[1] + 1] = Object.assign(e.sensors[event.data[1] + 1] || {}, { link: !0 });
            break;
        case 3:
            a = "Авария питания", d = "Уведомление", c = "Беспроводной датчик №" + (event.data[1] + 1), s = "kz.svg", p = "notification", e.sensors[event.data[1] + 1] = Object.assign(e.sensors[event.data[1] + 1] || {}, { battery: !0 });
            break;
        case 2:
            a = "Питание в норме", d = "Уведомление", c = "Беспроводной датчик №" + (event.data[1] + 1), s = "kz.svg", p = "notification", e.sensors[event.data[1] + 1] = Object.assign(e.sensors[event.data[1] + 1] || {}, { battery: !0 })
    }
});
