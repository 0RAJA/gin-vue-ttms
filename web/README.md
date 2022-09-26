ttms-web

node_modules\qrcodejs2文件源码有误，打包时把 if (this._android && this._android <= 2.1) 改为 if (this && this._android <= 2.1)