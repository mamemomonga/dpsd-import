# SIGMA DPM1,2,3,DPQ1 向け SDカードインポートツール

ファイル名を以下のように変換しながらコピーします。

* 105SIGMA/DP2M1234.JPG -> JPEG/DP2M_105-1234.JPG
* 106SIGMA/DP2M1234.X3F -> X3F/DP2M_106-1234.X3F

本プログラムはGolangで書かれています。macOS特有のファイル構造やコマンドに依存していますので、macOS専用です。また、ファイル名はDP??型で、sRGB保存である必要があります(AdobeRGBの場合はファイル名が変わります。)

初回起動時にインポート先フォルダを入力してください。こちらの情報は ~/.dpsd-import.yaml に保存されます。

機種名やマウント先フォルダは main.go にベタ書きしてあります。

# ビルド方法

    $ make

# 参考資料

* [CIPA DC-009-2010 カメラファイルシステム規格DCF2.0（2010年版）](https://www.cipa.jp/j/std/std-sec.html)
* [dp1 Quattro 取扱説明書](https://www.sigma-global.com/jp/cameras/dp1-quattro/)

# License

MIT