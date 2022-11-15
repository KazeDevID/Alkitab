# API Alkitab Indonesia


## [https://alkitab-kaze.herokuapp.com](https://alkitab-kaze.herokuapp.com)

API Alkitab Indonesia adalah sebuah *third-party backend* Alkitab yang sumber bacaannya berasal dari [sabda.org](https://sabda.org).

[sabda.org](https://sabda.org) adalah sebuah Yayasan Kristen non-profit dan non-komersial yang bergerak dalam bidang pelayanan media komputer dan Internet di Indonesia.

## Dokumentasi


### Passage Content V3

#### Base URL

Version | URL
-- | --
v3 | [https://alkitab-kaze.herokuapp.com/v3/passage/Yoh/1?ver=tb](https://alkitab-kaze.herokuapp.com/v3/passage/Yoh/1?ver=tb)

#### Example

Version | Method | URL
-- | -- | --
v3 | GET | [https://alkitab-kaze.herokuapp.com/v3/passage/{**chapter**}/{**chapter**}?ver={**ver**}](https://alkitab-kaze.herokuapp.com/v3/passage/Yoh/1?ver=tb)

#### Response

``` bash
{
    "title": "Yohanes 1:1-51",
    "book_number": 43,
    "chapter": 1,
    "verses": [
        {
            "verse": 0,
            "type": "title",
            "content": "Firman yang telah menjadi manusia"
        },
        {
            "verse": 1,
            "type": "text",
            "content": "Pada mulanya adalah Firman; Firman itu bersama-sama dengan Allah dan Firman itu adalah Allah."
        },
        ....
    ]
}
```

### Passage List

Untuk menampilkan passage list perjanjian lama dan baru

#### Example

Version | Method | URL
-- | -- | --
v2 | GET | [https://alkitab-kaze.herokuapp.com/v2/passage/list](https://alkitab-kaze.herokuapp.com/v2/passage/list)

#### Response

```json
{
    "passage_list": [
        {
            "book_number": 1,
            "abbreviation": "Kej",
            "book_name": "Kejadian",
            "total_chapter": 50
        },
        {
            "book_number": 2,
            "abbreviation": "Kel",
            "book_name": "Keluaran",
            "total_chapter": 40
        },
        ...
    ]
}
```

### Passage and Chapter

Untuk menampilkan kitab dan pasal Alkitab gunakan **passage/{passage}/{chapter}** dengan mengirim *request* nama kitab dan nomor pasal.

#### Base URL

Version | URL
-- | --
v2 | [https://alkitab-kaze.herokuapp.com/v2/passage/{**chapter**}/{**chapter**}?ver={**ver**}](https://alkitab-kaze.herokuapp.com/v2/passage/Yoh/1?ver=tb)
v1 | [https://alkitab-kaze.herokuapp.com/v1/passage/{**chapter**}/{**chapter**}](https://alkitab-kaze.herokuapp.com/v1/passage/Yoh/1)

Variabel | Keterangan | Tipe data | *Required*
-- | -- | -- | --
**passage** | Adalah kitab yang ingin ditampilkan. Menerima berbagai format dalam bahasa Indonesian dan Inggris (Yohanes, Yoh, John, dll). | `string` | Yes
**chapter** | Adalah nomor pasal yang ingin ditampilkan dari bagian Alkitab | `int` | Yes
**ver** | Adalah versi pilihan bahasa **tb** (terjemahan baru) *default*, **bis** (bahasa Indonesia sehari-hari), **net** (bahasa Inggris) | `string` | No

#### Example

Version | Method | URL
-- | -- | --
v2 | GET | [https://alkitab-kaze.herokuapp.com/v2/passage/Yoh/1?ver=tb](https://alkitab-kaze.herokuapp.com/v2/passage/Yoh/1?ver=tb)
v1 | GET | [https://alkitab-kaze.herokuapp.com/v1/passage/Yoh/1](https://alkitab-kaze.herokuapp.com/v1/passage/Yoh/1)

#### Response v2

```json
{
    "title": "Yohanes 1:1-51",
    "book_number": 43,
    "chapter": 1,
    "verses": [
        {
            "verse": 1,
            "content": "Pada mulanya adalah Firman; Firman itu bersama-sama dengan Allah dan Firman itu adalah Allah."
        },
        {
            "verse": 2,
            "content": "Ia pada mulanya bersama-sama dengan Allah."
        },
        ...
    ]
}
```

#### Response v1

```json
{
    "message": "success to fetching data",
    "response": {
        "title": "Yohanes 1:1-51",
        "passage": [
            {
                "book_id": 43,
                "title": "Yohanes 1:1-51",
                "chapter": {
                    "chap": 1,
                    "verses": {
                        "verse": [
                            {
                                "number": 1,
                                "text": "Pada mulanya adalah Firman; Firman itu bersama-sama dengan Allah dan Firman itu adalah Allah."
                            },
                            {
                                "number": 2,
                                "text": "Ia pada mulanya bersama-sama dengan Allah."
                            },
                            ...
                        ]
                    }
                }
            }
        ]
    },
    "status": 200
}
```

### Passage, Chapter, and Verse

Untuk membaca satu ayat Alkitab gunakan **passage/{passage}/{chapter}/{verse}** dengan mengirim nama kitab, nomor pasal dan nomor ayat Alkitab.

#### Base URL

Version | URL
-- | --
v2 | [https://alkitab-kaze.herokuapp.com/v2/passage/{**passage**}/{**chapter**}/{**verse**}?ver={**tb**}](https://alkitab-kaze.herokuapp.com/v2/passage/Yoh/1/1?ver=tb)
v1 | [https://alkitab-kaze.herokuapp.com/v1/passage/{**passage**}/{**chapter**}/{**verse**}](https://alkitab-kaze.herokuapp.com/v1/passage/Yoh/1/1)

Variabel | Keterangan | Tipe data | *Required*
-- | -- | -- | --
**passage** | Adalah kitab yang ingin ditampilkan. Menerima berbagai format dalam bahasa Indonesian dan Inggris (Yohanes, Yoh, John, dll). | `string` | Yes
**chapter** | Adalah nomor pasal yang ingin ditampilkan dari bagian Alkitab | `int` | Yes
**verse** | Adalah nomor ayat dari bagian pasal/bab yang ingin ditampilkan | `int` | Yes
**ver** | Adalah versi pilihan bahasa **tb** (terjemahan baru) *default*, **bis** (bahasa Indonesia sehari-hari), **net** (bahasa Inggris) | `string` | No

#### Example

Version | Method | URL
-- | -- | --
v2 | GET | [https://alkitab-kaze.herokuapp.com/v2/passage/Yoh/1/1?ver=tb](https://alkitab-kaze.herokuapp.com/v2/passage/Yoh/1/1?ver=tb)
v1 | GET | [https://alkitab-kaze.herokuapp.com/v1/passage/Yoh/1/1](https://alkitab-kaze.herokuapp.com/v1/passage/Yoh/1/1)

#### Response v2

```json
{
    "title": "Yohanes 1:1",
    "book_number": 43,
    "chapter": 1,
    "verses": [
        {
            "verse": 1,
            "content": "Pada mulanya adalah Firman; Firman itu bersama-sama dengan Allah dan Firman itu adalah Allah."
        }
    ]
}
```

#### Response v1

```json
{
    "message": "success to fetching data",
    "response": {
        "title": "Yohanes 1:1",
        "passage": [
            {
                "book_id": 43,
                "title": "Yohanes 1:1",
                "chapter": {
                    "chap": 1,
                    "verses": {
                        "verse": [
                            {
                                "number": 1,
                                "text": "Pada mulanya adalah Firman; Firman itu bersama-sama dengan Allah dan Firman itu adalah Allah."
                            }
                        ]
                    }
                }
            }
        ]
    },
    "status": 200
}
```

### Status Code

API mengembalikan beberapa status code sebagai berikut:

Status Code | Keterangan
-- | --
200 | OK
201 | CREATED
400 | BAD REQUEST
404 | NOT FOUND
500 | INTERNAL SERVER ERROR

### Author

API Alkitab Indonesia ini dibuat oleh **KazeDevID**. Untuk pertanyaan kritik dan saran, silahkan *drop* email ke [kazedevid@gmail.com](mailto:kazedevid@gmail.com).




---

> Kemuliaan bagi Tuhan saja - Tuhan Yesus memberkati 🤗
