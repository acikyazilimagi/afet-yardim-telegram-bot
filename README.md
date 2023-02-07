#afet-yardim-telegram-bot

###Bot adı ve kullanıcı adı
Telegram bot username: AfetYardim_bot <br>

Telegram bot name: AfetYardımBot <br>

###Address extract API
Gelen mesajdan adres çıkartmak için kullandığımız api: <br>
https://huggingface.co/spaces/mertcobanov/deprem-ocr-2
<br>

Not: API üzerinde performans geliştirmeleri devam ediyor. <br>

###MVP'de kullanılan regex
Eğer API'dan dönen veride şehir bilgisi boş ise
şehri regex kullanarak dolduruyoruz. Şimdilik yalnızca afet bölgesindeki şehirleri ve bu şehirlerin
yazımında oluşabilecek varyasyonları kullandık.

Regex: <br>
```
(?i)((gaz[ıiİI]antep)|(malatya)|(batman)|(b[ıiIİ]ng[oöOÖ]l)|(elaz[Iİıi][gğ])|(kilis)|(diyarbak[ıiIİ]r)|(mardin)|(siirt)|([SsŞş][ıiIİ]rnak)|(van)|(mu[sşSŞ])|(bitlis)|(hakkari)|(adana)|(osmaniye)|(hatay)|(kahramanmara[sşSŞ])|(mara[SŞsş])|(antep))
```