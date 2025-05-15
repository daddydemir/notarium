RFC: Kişisel Günlük & Not Yönetim Sistemi
RFC No: PKM-2024-001
Yazar: Claude
Oluşturma Tarihi: 15 Mayıs 2025
Durum: Taslak
İlgili Birimler: Frontend Ekibi, Backend Ekibi, UX Tasarım
## 1. Özet
Bu RFC, günlük bazlı kişisel not tutma ve bilgi yönetimi için tasarlanmış çok platformlu bir uygulama önerisi sunmaktadır. Sistem, kullanıcılara günlük kayıtlar oluşturma, konuları etiketleme, dosya ekleme ve çoklu cihaz senkronizasyonu özellikleri sunarak kişisel bilgi yönetimini (Personal Knowledge Management - PKM) kolaylaştırmayı amaçlamaktadır.
## 2. Motivasyon
Modern bilgi işçileri ve öğrenciler, giderek artan bir bilgi yığını ile baş etmek zorundadır. Mevcut not alma uygulamaları genellikle ya aşırı karmaşık ya da temel özelliklerden yoksundur. Bu proje, şu sorunları çözmeyi hedeflemektedir:

Günlük bazlı not tutma alışkanlığını teşvik etme
Bilgiyi etiketlerle anlamlı şekilde organize etme
Çoklu cihaz kullanımında kesintisiz deneyim sağlama
Markdown desteğiyle hızlı ve etkili formatlama imkanı sunma
Arama ve filtreleme ile bilgiye hızlı erişim

## 3. Tasarım Detayları

### 3.1 Sistem Mimarisi

```
┌─────────────┐      ┌─────────────┐      ┌─────────────┐
│   Frontend  │◄── ─►│    API      │◄────►│  PostgreSQL │
│  (Vanilla   │      │ (Golang)    │      │  Veritabanı │
│   JS/HTML)  │      │             │      │             │
└─────────────┘      └─────────────┘      └─────────────┘
```

### 3.2 Veri Modeli

**entries (Günlük Kayıtları)**
```
- id: UUID (PK, DEFAULT gen_random_uuid())
- date: DATE (NOT NULL, UNIQUE)
- title: VARCHAR(255)
- created_at: TIMESTAMPTZ (NOT NULL, DEFAULT now())
- updated_at: TIMESTAMPTZ
```

**topics (Konular)**
```
- id: UUID (PK, DEFAULT gen_random_uuid())
- entry_id: UUID (NOT NULL, FK -> entries.id, ON DELETE CASCADE)
- name: VARCHAR(100) (NOT NULL)
- is_starred: BOOLEAN (NOT NULL, DEFAULT false)
- is_completed: BOOLEAN (NOT NULL, DEFAULT false)
- order: INTEGER (DEFAULT 0)
- created_at: TIMESTAMPTZ (NOT NULL, DEFAULT now())
- updated_at: TIMESTAMPTZ
```

**notes (Notlar)**
```
- id: UUID (PK, DEFAULT gen_random_uuid())
- topic_id: UUID (NOT NULL, FK -> topics.id, ON DELETE CASCADE)
- content: TEXT (NOT NULL)
- version: INTEGER (NOT NULL, DEFAULT 1)
- created_at: TIMESTAMPTZ (NOT NULL, DEFAULT now())
- updated_at: TIMESTAMPTZ
```

**tags (Etiketler)**
```
- id: UUID (PK, DEFAULT gen_random_uuid())
- name: VARCHAR(50) (NOT NULL, UNIQUE)
- color: VARCHAR(7) (DEFAULT '#64748B')
- created_at: TIMESTAMPTZ (NOT NULL, DEFAULT now())
```

**topic_tags (Konu-Etiket İlişkisi)**
```
- topic_id: UUID (NOT NULL, FK -> topics.id, ON DELETE CASCADE)
- tag_id: UUID (NOT NULL, FK -> tags.id, ON DELETE CASCADE)
- created_at: TIMESTAMPTZ (NOT NULL, DEFAULT now())
- PRIMARY KEY (topic_id, tag_id)
```

**attachments (Ekler)**
```
- id: UUID (PK, DEFAULT gen_random_uuid())
- note_id: UUID (NOT NULL, FK -> notes.id, ON DELETE CASCADE)
- url: TEXT (NOT NULL)
- file_name: VARCHAR(255) (NOT NULL)
- mime_type: VARCHAR(100) (NOT NULL)
- size: BIGINT
- created_at: TIMESTAMPTZ (NOT NULL, DEFAULT now())
```

**reminders (Hatırlatıcılar)**
```
- id: UUID (PK, DEFAULT gen_random_uuid())
- topic_id: UUID (NOT NULL, FK -> topics.id, ON DELETE CASCADE)
- remind_at: TIMESTAMPTZ (NOT NULL)
- is_completed: BOOLEAN (NOT NULL, DEFAULT false)
- created_at: TIMESTAMPTZ (NOT NULL, DEFAULT now())
```

3.3 API Endpoints


## Günlük Kayıtları (Entries)

### Tüm Günlük Kayıtlarını Listele
**`GET /api/v1/entries`**

**Yanıt:**
```json
[
  {
    "id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77",
    "date": "2025-05-14",
    "title": "Yeni Başlangıç",
    "created_at": "2025-05-14T08:00:00Z"
  },
  {
    "id": "b1d5e1cd-4ec8-4f62-b48b-5464ad1bb8b5",
    "date": "2025-05-15",
    "title": "Bugünün Notları",
    "created_at": "2025-05-15T09:00:00Z"
  }
]
```

### Belirli Bir Tarihe Ait Günlük Kaydını Getir
**`GET /api/v1/entries/{date}`**

**İstek:**
```
GET /api/v1/entries/2025-05-14
```

**Yanıt:**
```json
{
  "id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77",
  "date": "2025-05-14",
  "title": "Yeni Başlangıç",
  "created_at": "2025-05-14T08:00:00Z"
}
```

### Yeni Günlük Kaydı Oluştur
**`POST /api/v1/entries`**

**İstek:**
```json
{
  "date": "2025-05-15",
  "title": "Bugünün Notları"
}
```

**Yanıt:**
```json
{
  "id": "b1d5e1cd-4ec8-4f62-b48b-5464ad1bb8b5",
  "date": "2025-05-15",
  "title": "Bugünün Notları",
  "created_at": "2025-05-15T09:00:00Z"
}
```

### Günlük Kaydını Sil
**`DELETE /api/v1/entries/{id}`**

**İstek:**
```
DELETE /api/v1/entries/e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77
```

**Yanıt:**
```json
{
  "status": "deleted"
}
```

## Konular (Topics)

### Tüm Konuları Listele
**`GET /api/v1/topics`**

**Yanıt:**
```json
[
  {
    "id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
    "entry_id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77",
    "name": "Go'da Interface Kullanımı",
    "is_starred": false,
    "is_completed": false,
    "order": 0,
    "created_at": "2025-05-15T09:15:00Z",
    "updated_at": null
  },
  {
    "id": "f4a9dd8c-3f41-5090-0dd1-5c8e4903e0d2",
    "entry_id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77",
    "name": "API Tasarımı",
    "is_starred": true,
    "is_completed": false,
    "order": 1,
    "created_at": "2025-05-15T10:00:00Z",
    "updated_at": null
  }
]
```

### Belirli Bir Konuyu Getir
**`GET /api/v1/topics/{id}`**

**İstek:**
```
GET /api/v1/topics/d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1
```

**Yanıt:**
```json
{
  "id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
  "entry_id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77",
  "name": "Go'da Interface Kullanımı",
  "is_starred": false,
  "is_completed": false,
  "order": 0,
  "created_at": "2025-05-15T09:15:00Z",
  "updated_at": null
}
```

### Belirli Bir Günlük Kaydına Ait Konuları Getir
**`GET /api/v1/topics?entry_id={entry_id}`**

**İstek:**
```
GET /api/v1/topics?entry_id=e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77
```

**Yanıt:**
```json
[
  {
    "id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
    "entry_id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77",
    "name": "Go'da Interface Kullanımı",
    "is_starred": false,
    "is_completed": false,
    "order": 0,
    "created_at": "2025-05-15T09:15:00Z",
    "updated_at": null
  }
]
```

### Belirli Bir Etikete Sahip Konuları Getir
**`GET /api/v1/topics?tag={tag_name}`**

**İstek:**
```
GET /api/v1/topics?tag=backend
```

**Yanıt:**
```json
[
  {
    "id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
    "entry_id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77",
    "name": "Go'da Interface Kullanımı",
    "is_starred": false,
    "is_completed": false,
    "order": 0,
    "created_at": "2025-05-15T09:15:00Z",
    "updated_at": null
  }
]
```

### Yeni Konu Oluştur
**`POST /api/v1/topics`**

**İstek:**
```json
{
  "entry_id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77",
  "name": "Go'da Interface Kullanımı",
  "order": 0
}
```

**Yanıt:**
```json
{
  "id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
  "entry_id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77",
  "name": "Go'da Interface Kullanımı",
  "is_starred": false,
  "is_completed": false,
  "order": 0,
  "created_at": "2025-05-15T09:15:00Z",
  "updated_at": null
}
```

### Konuyu Güncelle
**`PUT /api/v1/topics/{id}`**

**İstek:**
```json
{
  "name": "Go'da Interface Kullanımı - Güncellenmiş",
  "is_starred": true,
  "is_completed": false,
  "order": 1
}
```

**Yanıt:**
```json
{
  "id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
  "entry_id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77",
  "name": "Go'da Interface Kullanımı - Güncellenmiş",
  "is_starred": true,
  "is_completed": false,
  "order": 1,
  "created_at": "2025-05-15T09:15:00Z",
  "updated_at": "2025-05-15T10:00:00Z"
}
```

### Konuyu Sil
**`DELETE /api/v1/topics/{id}`**

**İstek:**
```
DELETE /api/v1/topics/d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1
```

**Yanıt:**
```json
{
  "status": "deleted"
}
```

### Konuyu Yıldızla/Yıldızı Kaldır
**`PUT /api/v1/topics/{id}/star`**

**İstek:**
```json
{
  "is_starred": true
}
```

**Yanıt:**
```json
{
  "id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
  "is_starred": true
}
```

## Notlar (Notes)

### Belirli Bir Konuya Ait Notları Getir
**`GET /api/v1/notes?topic_id={topic_id}`**

**İstek:**
```
GET /api/v1/notes?topic_id=d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1
```

**Yanıt:**
```json
[
  {
    "id": "8e6a7139-6c23-4f7e-95db-bdfefc5efab4",
    "topic_id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
    "content": "Go'da interface'ler davranışı tanımlar.",
    "version": 1,
    "created_at": "2025-05-15T09:20:00Z",
    "updated_at": null
  }
]
```

### Belirli Bir Notu Getir
**`GET /api/v1/notes/{id}`**

**İstek:**
```
GET /api/v1/notes/8e6a7139-6c23-4f7e-95db-bdfefc5efab4
```

**Yanıt:**
```json
{
  "id": "8e6a7139-6c23-4f7e-95db-bdfefc5efab4",
  "topic_id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
  "content": "Go'da interface'ler davranışı tanımlar.",
  "version": 1,
  "created_at": "2025-05-15T09:20:00Z",
  "updated_at": null
}
```

### Yeni Not Oluştur
**`POST /api/v1/notes`**

**İstek:**
```json
{
  "topic_id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
  "content": "Go'da interface'ler davranışı tanımlar."
}
```

**Yanıt:**
```json
{
  "id": "8e6a7139-6c23-4f7e-95db-bdfefc5efab4",
  "topic_id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
  "content": "Go'da interface'ler davranışı tanımlar.",
  "version": 1,
  "created_at": "2025-05-15T09:20:00Z",
  "updated_at": null
}
```

### Notu Güncelle
**`PUT /api/v1/notes/{id}`**

**İstek:**
```json
{
  "content": "Go'da interface'ler davranışı tanımlar ve tip güvenliği sağlar."
}
```

**Yanıt:**
```json
{
  "id": "8e6a7139-6c23-4f7e-95db-bdfefc5efab4",
  "topic_id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
  "content": "Go'da interface'ler davranışı tanımlar ve tip güvenliği sağlar.",
  "version": 2,
  "created_at": "2025-05-15T09:20:00Z",
  "updated_at": "2025-05-15T10:00:00Z"
}
```

### Notu Sil
**`DELETE /api/v1/notes/{id}`**

**İstek:**
```
DELETE /api/v1/notes/8e6a7139-6c23-4f7e-95db-bdfefc5efab4
```

**Yanıt:**
```json
{
  "status": "deleted"
}
```

## Etiketler (Tags)

### Tüm Etiketleri Listele
**`GET /api/v1/tags`**

**Yanıt:**
```json
[
  {
    "id": "a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6",
    "name": "backend",
    "color": "#FF5733",
    "created_at": "2025-05-15T09:25:00Z"
  },
  {
    "id": "b2c3d4e5-f6g7-8h9i-0j1k-l2m3n4o5p6q7",
    "name": "frontend",
    "color": "#33FF57",
    "created_at": "2025-05-15T09:30:00Z"
  }
]
```

### Yeni Etiket Oluştur
**`POST /api/v1/tags`**

**İstek:**
```json
{
  "name": "backend",
  "color": "#FF5733"
}
```

**Yanıt:**
```json
{
  "id": "a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6",
  "name": "backend",
  "color": "#FF5733",
  "created_at": "2025-05-15T09:25:00Z"
}
```

### Etiketi Sil
**`DELETE /api/v1/tags/{id}`**

**İstek:**
```
DELETE /api/v1/tags/a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6
```

**Yanıt:**
```json
{
  "status": "deleted"
}
```

### Konuya Etiket Ekle
**`POST /api/v1/topics/{topic_id}/tags/{tag_id}`**

**İstek:**
```
POST /api/v1/topics/d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1/tags/a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6
```

**Yanıt:**
```json
{
  "topic_id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
  "tag_id": "a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6",
  "created_at": "2025-05-15T09:35:00Z"
}
```

### Konudan Etiket Kaldır
**`DELETE /api/v1/topics/{topic_id}/tags/{tag_id}`**

**İstek:**
```
DELETE /api/v1/topics/d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1/tags/a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6
```

**Yanıt:**
```json
{
  "status": "deleted"
}
```

## Ekler (Attachments)

### Belirli Bir Nota Ait Ekleri Listele
**`GET /api/v1/attachments?note_id={note_id}`**

**İstek:**
```
GET /api/v1/attachments?note_id=8e6a7139-6c23-4f7e-95db-bdfefc5efab4
```

**Yanıt:**
```json
[
  {
    "id": "fa123456-7890-4bcd-9876-abcdef123456",
    "note_id": "8e6a7139-6c23-4f7e-95db-bdfefc5efab4",
    "url": "/uploads/slides.pdf",
    "file_name": "slides.pdf",
    "mime_type": "application/pdf",
    "size": 5242880,
    "created_at": "2025-05-15T09:30:00Z"
  }
]
```

### Yeni Ek Oluştur (Dosya Yükleme)
**`POST /api/v1/attachments`**

**İstek (multipart/form-data):**
```
{
  "note_id": "8e6a7139-6c23-4f7e-95db-bdfefc5efab4",
  "file": "slides.pdf"
}
```

**Yanıt:**
```json
{
  "id": "fa123456-7890-4bcd-9876-abcdef123456",
  "note_id": "8e6a7139-6c23-4f7e-95db-bdfefc5efab4",
  "url": "/uploads/slides.pdf",
  "file_name": "slides.pdf",
  "mime_type": "application/pdf",
  "size": 5242880,
  "created_at": "2025-05-15T09:30:00Z"
}
```

### Belirli Bir Eki İndir
**`GET /api/v1/attachments/{id}`**

**İstek:**
```
GET /api/v1/attachments/fa123456-7890-4bcd-9876-abcdef123456
```

**Yanıt:**
```
// Binary file stream (e.g., application/pdf)
```

### Eki Sil
**`DELETE /api/v1/attachments/{id}`**

**İstek:**
```
DELETE /api/v1/attachments/fa123456-7890-4bcd-9876-abcdef123456
```

**Yanıt:**
```json
{
  "status": "deleted"
}
```

## Arama (Search)

### Metin İçeriklerinde Arama Yap
**`GET /api/v1/search?q={query}`**

**İstek:**
```
GET /api/v1/search?q=interface
```

**Yanıt:**
```json
[
  {
    "type": "note",
    "content": "Go'da interface'ler davranışı tanımlar.",
    "note_id": "8e6a7139-6c23-4f7e-95db-bdfefc5efab4",
    "topic_id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
    "entry_id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77"
  }
]
```

### Belirli Etiketlere Göre Arama Yap
**`GET /api/v1/search?tags={tag1,tag2}`**

**İstek:**
```
GET /api/v1/search?tags=backend,frontend
```

**Yanıt:**
```json
[
  {
    "type": "topic",
    "id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
    "name": "Go'da Interface Kullanımı",
    "entry_id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77",
    "tags": ["backend"]
  }
]
```

## İçe/Dışa Aktarma (Import/Export)

### Tüm Verileri JSON Formatında Dışa Aktar
**`GET /api/v1/export`**

**Yanıt:**
```json
{
  "entries": [
    {
      "id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77",
      "date": "2025-05-14",
      "title": "Yeni Başlangıç",
      "created_at": "2025-05-14T08:00:00Z"
    }
  ],
  "topics": [
    {
      "id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
      "entry_id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77",
      "name": "Go'da Interface Kullanımı",
      "is_starred": false,
      "is_completed": false,
      "order": 0,
      "created_at": "2025-05-15T09:15:00Z",
      "updated_at": null
    }
  ],
  "notes": [
    {
      "id": "8e6a7139-6c23-4f7e-95db-bdfefc5efab4",
      "topic_id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
      "content": "Go'da interface'ler davranışı tanımlar.",
      "version": 1,
      "created_at": "2025-05-15T09:20:00Z",
      "updated_at": null
    }
  ],
  "tags": [
    {
      "id": "a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6",
      "name": "backend",
      "color": "#FF5733",
      "created_at": "2025-05-15T09:25:00Z"
    }
  ]
}
```

### JSON Verilerini İçe Aktar
**`POST /api/v1/import`**

**İstek:**
```json
{
  "entries": [
    {
      "date": "2025-05-15",
      "title": "Bugünün Notları"
    }
  ],
  "topics": [
    {
      "entry_id": "e2c9a246-b56a-4b6b-8b76-6a1fc2c14a77",
      "name": "Go'da Interface Kullanımı"
    }
  ],
  "notes": [
    {
      "topic_id": "d3b8cc7b-2f30-4989-9cc0-4b7d3892d9c1",
      "content": "Go'da interface'ler davranışı tanımlar."
    }
  ],
  "tags": [
    {
      "name": "backend",
      "color": "#FF5733"
    }
  ]
}
```

**Yanıt:**
```json
{
  "status": "imported"
}
```
## 3.4 Frontend Yapısı

Frontend mimarisinde MVC (Model-View-Controller) desenine benzer bir yapı kullanılacak:

### Sayfalar
- Ana sayfa (Günlük görünümü)
- Konu detay sayfası
- Etiket bazlı filtreleme sayfası
- Arama sonuçları sayfası
- Ayarlar sayfası
- İçe/Dışa aktarma sayfası

### Bileşenler
- Markdown editörü
- Dosya yükleme bileşeni
- Etiket seçici
- Takvim widget'ı
- Hatırlatıcı oluşturucu

Frontend, responsive tasarım prensiplerine uygun olarak geliştirilecek ve mobil cihazlarda da sorunsuz çalışacaktır.

## 4. Teknik Zorluklar ve Çözümler

### 4.1 Markdown İşleme
Markdown metnini HTML'e dönüştürmek için client-side bir Markdown parser kütüphanesi kullanılacak. Önerilen kütüphaneler:
- marked.js
- showdown.js

### 4.2 Veri Senkronizasyonu
Çoklu cihaz senkronizasyonu için:
- Her işlem sonrası API ile iletişim kurulacak
- Çevrimdışı kullanım için IndexedDB kullanılacak
- Çevrimdışı değişiklikler, çevrimiçi olunduğunda otomatik senkronize edilecek
- Çakışma çözümü için "son yazma kazanır" politikası kullanılacak

### 4.3 Dosya Depolama
Dosyalar için iki seçenek değerlendirilecek:
- Yerel dosya sistemi (daha basit ve hızlı başlangıç için)
- S3-uyumlu bir nesne depolama hizmeti (ölçeklenebilirlik için)

Başlangıçta 1. seçeneğin uygulanması, daha sonra gerekirse 2. seçeneğe geçiş planlanmaktadır.

## 5. Uygulama Takvimi

| Aşama                     | Süre       | Açıklama                                                                 |
|---------------------------|------------|--------------------------------------------------------------------------|
| Tasarım ve Planlama       | 2 hafta    | Teknik tasarım, UI/UX tasarımı finalize edilecek                         |
| Veritabanı Kurulumu       | 1 hafta    | Şema oluşturma, migration yazılması                                      |
| Backend API Geliştirme    | 3 hafta    | Temel API endpoint'lerinin uygulanması                                   |
| Frontend Geliştirme       | 4 hafta    | Ana sayfa, konu sayfası ve temel işlevlerin geliştirilmesi               |
| Entegrasyon               | 2 hafta    | Frontend ve backend entegrasyonu                                        |
| Test                      | 2 hafta    | Birim ve entegrasyon testleri                                           |
| Beta Sürümü               | 1 hafta    | İlk kullanıcı testleri ve geri bildirim alma                             |
| İyileştirme               | 2 hafta    | Beta geri bildirimlerine dayalı iyileştirmeler                          |
| İlk Sürüm                 | -          | MVP özellikleriyle ilk kararlı sürüm                                    |

**Toplam geliştirme süresi:** 15 hafta

## 6. Ölçme ve Başarı Kriterleri

Projenin başarısı aşağıdaki metriklerle değerlendirilecektir:

### Kullanıcı Deneyimi Metrikleri
- Ortalama günlük aktif kullanıcı sayısı
- Ortalama oturum süresi
- Günlük oluşturulan not sayısı

### Teknik Metrikler
- API yanıt süreleri
- Veritabanı sorgu performansı
- Frontend yükleme süresi

### Başarı Kriterleri
- Hatasız çoklu cihaz senkronizasyonu
- 1 saniyeden kısa API yanıt süreleri
- 100 MB'a kadar dosya desteği

## 7. Gelecek Aşama Özellikleri

Şu an kapsam dışında bırakılan ancak gelecekte eklenebilecek özellikler:

### Kullanıcı Yönetimi ve Yetkilendirme
- Oturum açma/kayıt olma
- Rol tabanlı erişim kontrolü
- Not paylaşımı özellikleri

### Gelişmiş Özellikler
- Kanban görünümü
- AI tabanlı içerik önerileri
- Otomatik etiketleme
- OCR ile görsel metin tanıma

### Entegrasyonlar
- Takvim entegrasyonu
- Email entegrasyonu
- Slack/Discord entegrasyonu

## 8. Alternatifler ve Değerlendirme

Bu bölümde, alternatif teknoloji teklifleri ve tasarım kararları değerlendirilmektedir:

### 8.1 Frontend Alternatifi: React
**Avantajlar:**
- Büyük ekosistem ve hazır bileşenler
- Daha yapılandırılmış kod organizasyonu

**Dezavantajlar:**
- Daha fazla bağımlılık ve karmaşıklık
- Daha uzun derleme süresi

### 8.2 Backend Alternatifi: Node.js/Express
**Avantajlar:**
- JavaScript'in her yerde kullanılması
- Zengin NPM ekosistemi

**Dezavantajlar:**
- Go'ya göre daha düşük performans
- Daha yüksek kaynak kullanımı

### 8.3 Veritabanı Alternatifi: MongoDB
**Avantajlar:**
- Şemasız yapı ile esnek geliştirme
- JSON benzeri doküman yapısı

**Dezavantajlar:**
- İlişkisel sorgular için zorluklar
- Veri tutarlılığı riski

Mevcut seçimlerin (Vanilla JS, Go, PostgreSQL) daha hızlı geliştirme ve daha az karmaşıklık sunması nedeniyle tercih edildiği değerlendirilmiştir.

## 9. Güvenlik Değerlendirmesi

Bu projenin güvenlik açısından dikkat edilmesi gereken noktalar:

### Veri Güvenliği
- Kişisel notların güvenli depolanması
- Güvenli dosya yükleme ve doğrulama

### API Güvenliği
- Rate limiting uygulanması
- İleriki aşamalarda JWT tabanlı kimlik doğrulama eklenecek

### İstemci Güvenliği
- XSS saldırılarına karşı markdown sanitizasyonu
- CSRF koruması

Güvenlik açıkları için düzenli kod incelemeleri yapılacak ve gerektiğinde dış güvenlik değerlendirmeleri alınacaktır.

## 10. Sonuç

Bu RFC, basit ancak güçlü bir kişisel bilgi yönetim sistemi önerir. Temel odak noktası kullanıcı deneyimi, verimlilik ve çoklu cihaz desteğidir. Teknoloji seçimleri, hızlı geliştirme ve bakım kolaylığı dengesini gözetmektedir.

### Referanslar
- [PostgreSQL Dokümantasyonu](https://www.postgresql.org/docs/)
- [Go Web Geliştirme](https://golang.org/doc/)
- [Markdown Spesifikasyonu](https://commonmark.org/)
- [API Tasarım Prensipleri](https://restfulapi.net/)