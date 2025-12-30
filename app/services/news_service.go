package services

import (
	"context"
	"donasitamanzakattest/app/models"
	"donasitamanzakattest/app/repositories"
	"donasitamanzakattest/app/web"
	"donasitamanzakattest/config"
	"donasitamanzakattest/pkg/util"
	"errors"
	"sync"
	"time"
)

type NewsStore struct {
	mu    sync.RWMutex
	items map[string]*models.News
	order []string
}

var (
	newsStoreOnce sync.Once
	newsStore     *NewsStore
)

func GetNewsStore() *NewsStore {
	newsStoreOnce.Do(func() {
		store := &NewsStore{
			items: map[string]*models.News{},
			order: []string{},
		}

		seedNewss := []*models.News{
			{
				Id:       store.generateUniqueID(12),
				Owner:    "Aksi Taza",
				HeadLine: "Hadirkan Keceriaan Yatim, Taman Zakat Ajak Fun Trip dan Beri Perlengkapan Sekolah",
				Lead:     "Hadirkan Keceriaan Yatim, Taman Zakat Ajak Fun Trip dan Beri Perlengkapan Sekolah",
				Body: `<h2>Semarak Muharram: Taman Zakat Gelar Fun Trip dan Bagi Perlengkapan Sekolah untuk Anak Yatim</h2>

<p>
Dalam rangka menyambut tahun ajaran baru sekaligus menyemarakkan bulan Muharram, 
<strong>Lembaga Amil Zakat Taman Zakat</strong> menghadirkan program penuh keceriaan 
bagi anak-anak yatim melalui kegiatan <em>Fun Trip</em> edukatif dan pembagian 
paket perlengkapan sekolah.
</p>

<p>
Sebanyak <strong>15 anak yatim</strong> dari Panti Asuhan Muzdalifah, Waru, Sidoarjo, 
diajak menikmati pengalaman tak terlupakan dengan mengunjungi 
<strong>Museum Penerbangan TNI AL Juanda</strong> dan 
<strong>Rumah Pintar Cendekia Juanda</strong> pada Kamis (10/7/2025).
</p>

<h3>Pembagian Perlengkapan Sekolah</h3>

<p>
Kegiatan dimulai sejak pukul 08.00 WIB. Sebelum petualangan dimulai, 
anak-anak menerima paket perlengkapan sekolah berupa <strong>tas, buku, dan alat tulis</strong> 
sebagai bekal menyambut tahun ajaran baru.
</p>

<p>
“<em>Ini bentuk dukungan kami agar anak-anak yatim bisa memulai sekolah dengan semangat baru 
dan hati yang bahagia. Kami ingin mereka tahu bahwa mereka tidak sendiri</em>,” 
ujar <strong>Ziyad</strong>, General Manager Taman Zakat.
</p>

<h3>Kunjungan ke Museum Penerbangan</h3>

<p>
Dengan penuh semangat, anak-anak menaiki kereta kelinci menuju Museum Penerbangan. 
Di sana mereka menjelajahi kabin pesawat, melihat proses pendaratan pesawat secara langsung, 
hingga mengenal sejarah dunia aviasi.
</p>

<p>
“<em>Harapannya, anak-anak mendapatkan pengalaman yang bisa membangkitkan rasa ingin tahu 
dan cita-cita mereka di masa depan</em>,” tambah Ziyad.
</p>

<h3>Aktivitas Seru di Rumah Pintar Cendekia</h3>

<p>
Setelah dari museum, kegiatan berlanjut ke Rumah Pintar Cendekia Juanda. 
Di sini, anak-anak mengikuti berbagai aktivitas menyenangkan seperti 
<strong>memberi makan burung dan kura-kura, bermain flying fox, senam bersama, 
menonton film edukasi, dan mewarnai</strong> di aula.
</p>

<p>
Seluruh rangkaian kegiatan didampingi oleh <strong>Umi Muzdalifah</strong>, 
Kepala Panti Asuhan Muzdalifah, serta dua relawan mahasiswa Unesa, yaitu 
<strong>Kak Nita</strong> dan <strong>Kak Shinta</strong>.
</p>

<p>
“<em>Anak-anak sangat senang. Bagi mereka, hari ini bukan hanya perjalanan biasa, 
tapi momen yang akan terus mereka kenang</em>,” ujar Umi Muzdalifah.
</p>

<h3>Penutup Penuh Kebahagiaan</h3>

<p>
Acara ditutup dengan makan siang bersama, menciptakan suasana hangat dan penuh kebersamaan. 
Tawa riang dan senyum bahagia anak-anak menjadi bukti bahwa hari itu benar-benar bermakna.
</p>

<p>
Melalui kegiatan ini, Taman Zakat kembali menegaskan komitmennya untuk tidak hanya menyalurkan 
zakat secara materi, tetapi juga menghadirkan program-program yang menyentuh hati dan 
membangun masa depan.
</p>

<p>
“<em>InsyaAllah, ini bukan yang terakhir. Kami ingin terus hadir, memberi manfaat, 
dan menyalakan harapan bagi mereka yang membutuhkan</em>,” pungkas Ziyad.
</p>
`,
				Tail:      "",
				CreatedAt: time.Now().Format(time.DateTime),
				Image:     `https://donasi.tamanzakat.org/wp-content/uploads/2025/07/Taman-Zakat-Ajak-Fun-Trip-dan-Beri-Perlengkapan-Sekolah.jpg`,
			},
		}

		for _, news := range seedNewss {
			store.items[news.Id] = news
			store.order = append(store.order, news.Id)
		}

		newsStore = store
	})

	return newsStore
}

func (s *NewsStore) generateUniqueID(length int) string {
	for {
		id := util.GenerateRandomString(length)
		if _, exists := s.items[id]; !exists {
			return id
		}
	}
}

func (s *NewsStore) Create(news *web.NewsUpsertRequest) *models.News {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.generateUniqueID(12)
	cat := &models.News{
		Id:        id,
		Owner:     news.Owner,
		HeadLine:  news.HeadLine,
		Lead:      news.Lead,
		Body:      news.Body,
		Tail:      news.Tail,
		CreatedAt: time.Now().Format(time.DateTime),
		Image:     news.Image,
	}
	s.items[id] = cat
	s.order = append(s.order, id)
	return cat
}

func (s *NewsStore) List() []*models.News {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]*models.News, 0, len(s.order))
	for _, id := range s.order {
		if c, ok := s.items[id]; ok {
			out = append(out, &models.News{
				Id:        c.Id,
				Owner:     c.Owner,
				HeadLine:  c.HeadLine,
				Lead:      c.Lead,
				Body:      c.Body,
				Tail:      c.Tail,
				CreatedAt: c.CreatedAt,
				Image:     c.Image,
			})
		}
	}
	return out
}

func (s *NewsStore) GetByID(id string) (*models.News, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	c, ok := s.items[id]
	if !ok {
		return nil, false
	}
	return &models.News{
		Id:        c.Id,
		Owner:     c.Owner,
		HeadLine:  c.HeadLine,
		Lead:      c.Lead,
		Body:      c.Body,
		Tail:      c.Tail,
		CreatedAt: c.CreatedAt,
		Image:     c.Image,
	}, true
}

func (s *NewsStore) Update(id string, news *models.News) (*models.News, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, ok := s.items[news.Id]
	if !ok {
		return nil, false
	}
	c.Owner = news.Owner
	c.HeadLine = news.HeadLine
	c.Lead = news.Lead
	c.Body = news.Body
	c.Tail = news.Tail
	c.Image = news.Image
	return &models.News{
		Id:        c.Id,
		Owner:     c.Owner,
		HeadLine:  c.HeadLine,
		Lead:      c.Lead,
		Body:      c.Body,
		Tail:      c.Tail,
		CreatedAt: c.CreatedAt,
		Image:     c.Image,
	}, true
}

func (s *NewsStore) Delete(id string) (*models.News, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, ok := s.items[id]
	if !ok {
		return nil, false
	}

	deleted := &models.News{
		Id:        c.Id,
		Owner:     c.Owner,
		HeadLine:  c.HeadLine,
		Lead:      c.Lead,
		Body:      c.Body,
		Tail:      c.Tail,
		CreatedAt: c.CreatedAt,
		Image:     c.Image,
	}
	delete(s.items, id)

	for i := 0; i < len(s.order); i++ {
		if s.order[i] == id {
			s.order = append(s.order[:i], s.order[i+1:]...)
			break
		}
	}

	return deleted, true
}

type NewsService struct {
	service *Service
	store   *NewsStore
}

func NewNewsService(ctx context.Context, r *repositories.RepositoryContext, cfg *config.Config) *NewsService {
	return &NewsService{
		service: NewService(ctx, r, cfg),
		store:   GetNewsStore(),
	}
}

func (s *NewsService) Create(payload *web.NewsUpsertRequest) (*models.News, error) {

	return s.store.Create(payload), nil
}

func (s *NewsService) List() ([]*models.News, error) {
	return s.store.List(), nil
}

func (s *NewsService) GetByID(id string) (*models.News, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	c, ok := s.store.GetByID(id)
	if !ok {
		return nil, errors.New("news not found")
	}
	return c, nil
}

func (s *NewsService) Update(id string, payload *models.News) (*models.News, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	c, ok := s.store.Update(id, payload)
	if !ok {
		return nil, errors.New("news not found")
	}
	return c, nil
}

func (s *NewsService) Delete(id string) (*models.News, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	c, ok := s.store.Delete(id)
	if !ok {
		return nil, errors.New("news not found")
	}
	return c, nil
}
