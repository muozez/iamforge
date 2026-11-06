# ADR-0001: Saf Go Python Parser Seçimi

## Durum
Kabul Edildi (Accepted) – 2026-11-07

## Bağlam
`iamforge`, AWS Lambda ve altyapı kodlarını yerelde statik analizle tarayıp least-privilege IAM politikaları üreten bir CLI aracıdır. Projenin temel ilkelerinden biri:
- CGO bağımlılığı olmadan (`CGO_ENABLED=0`),
- Tek bir bağımsız statik binary olarak,
- Tüm hedef işletim sistemlerinde (Linux, macOS, Windows; amd64 ve arm64) sorunsuz derlenebilmesidir.

Geleneksel `tree-sitter` Go bağlayıcıları (örn. `smacker/go-tree-sitter` veya resmi repo) CGO gerektirmekte ve cross-compilation süreçlerini zorlaştırmaktadır. Bu nedenle Faz 0 spike çalışmasında CGO gerektirmeyen saf Go (pure-Go) alternatifler değerlendirilmiştir.

### Değerlendirilen Adaylar
1. **`github.com/odvcencio/gotreesitter`**: Tree-sitter çalışma zamanının ve gramerlerinin CGO'suz, saf Go ile yeniden yazılmış hali.
2. **`wazero` + Tree-sitter WASM**: Tree-sitter C kodunun WebAssembly'ye derlenip pure-Go WASM çalışma zamanı üzerinden koşturulması.
3. **`gpython` parser**: Python 3.4 gramerine dayalı saf Go parser.

## Karar
Python AST/CST çıkarımı için **`github.com/odvcencio/gotreesitter`** ve `github.com/odvcencio/gotreesitter/grammars/python` seçilmiştir.

Parser doğrudan kod tabanına bağlanmamış; `internal/extract/parser.go` altında tanımlanan `Parser`, `Tree` ve `Node` arayüzleri ile soyutlanmıştır.

## Gerekçe ve Spike Sonuçları
`internal/extract/python` altında 5 farklı modern Python sözdizimi içeren test dosyası ile doğrulama yapılmıştır:

| Kriter / Sözdizimi | Test Dosyası | Sonuç |
|---|---|---|
| f-string interpolasyonu | `spike_fstring.py` | Başarılı (`interpolation`) |
| Fonksiyon dekoratörleri | `spike_decorator.py` | Başarılı (`decorated_definition`) |
| Tip ipuçları (type hints) | `spike_typehint.py` | Başarılı (`type`) |
| Asenkron fonksiyonlar (`async/await`) | `spike_async.py` | Başarılı (`function_definition`) |
| Pattern Matching (Python 3.10+ `match`/`case`) | `spike_match.py` | Başarılı (`match_statement`) |

Spike doğrulamaları:
- Tüm testler `CGO_ENABLED=0` ile hatasız geçmiş (`HasError() == false`).
- Linux, Darwin ve Windows ortamları için çapraz derleme doğrulanmıştır.
- `gpython`'ın aksine modern Python sözdizimleri (3.10+ pattern matching, f-strings, modern type annotations) tam desteklenmektedir.
- `wazero` yaklaşımına kıyasla WASM modül yükleme ve bellek köprüsü karmaşıklığı olmadan doğal Go veri yapıları sunmaktadır.

## Riskler ve Azaltma Yolları
- **Gramer Güncellemeleri:** Python diline gelebilecek yeni sözdizimi eklemeleri `gotreesitter` kütüphanesinin güncellenmesiyle takip edilecektir.
- **Backend Değişikliği:** İleride parser değiştirilmek istenirse, `internal/extract.Parser` arayüzü sayesinde sadece `internal/extract/python` paketi etkilenecek; üst katmanlardaki extractor, mapping ve policy mantığı değişmeyecektir.
