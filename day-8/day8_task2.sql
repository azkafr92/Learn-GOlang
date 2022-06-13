-- INSERT
-- 1a. Insert 5 operators pada table operators
INSERT INTO `operators` (`name`, `created_at`, `updated_at`)
VALUES ('Telkomsel', NOW(), NOW()),
    ('Smartfren', NOW(), NOW()),
    ('XL', NOW(), NOW()),
    ('Indosat', NOW(), NOW()),
    ('Hinet', NOW(), NOW());
-- 1b. Insert 3 product type
INSERT INTO `product_types` (`name`, `created_at`, `updated_at`)
VALUES ('Pascabayar', NOW(), NOW()),
    ('Prabayar', NOW(), NOW()),
    ('Digital', NOW(), NOW());
-- 1c. Insert 2 product dengan product type id = 1 (Pascabayar) dan operators id = 3 (XL)
INSERT INTO `products` (
        `name`,
        `product_type_id`,
        `operator_id`,
        `created_at`,
        `updated_at`
    )
VALUES ('Silver', 1, 3, NOW(), NOW()),
    ('Gold', 1, 3, NOW(), NOW());
-- 1d. Insert 3 product dengan product type id = 2 (Prabayar) dan operators id = 1 (Telkomsel)
INSERT INTO `products` (
        `name`,
        `product_type_id`,
        `operator_id`,
        `created_at`,
        `updated_at`
    )
VALUES ('simPATI', 2, 1),
    ('KARTU As', 2, 1),
    ('LOOP', 2, 1);
-- 1e. Insert 3 product dengan product type id = 3 (Digital) dan operators id = 4 (Indosat)
INSERT INTO `products` (
        `name`,
        `product_type_id`,
        `operator_id`,
        `created_at`,
        `updated_at`
    )
VALUES ('Freedom Internet', 3, 4),
    ('Freedom U', 3, 4),
    ('Freedom Combo', 3, 4);
-- 1f. Insert product description pada setiap product
INSERT INTO `product_descriptions` (`id`, `description`, `created_at`, `updated_at`)
VALUES (
        1,
        'Unlimited Nelp ke XL & AXIS\n100 Menit Nelp ke Operator Lain\nUnlimited SMS AnyNet™',
        NOW(),
        NOW()
    ),
    (
        2,
        'Unlimited Nelp ke XL & AXIS\n150 Menit Nelp ke Operator Lain\nUnlimited SMS AnyNet™, Total FUP 40GB',
        NOW(),
        NOW()
    ),
    (
        3,
        'Simpati (nama digayakan sebagai simPATI) adalah produk dari Telkomsel. Kartu ini diluncurkan pada tahun 1997. Pada tanggal 18 Juni 2021, kartu ini dileburkan bersama Kartu As dan LOOP menjadi Telkomsel PraBayar dengan diluncurkannya logo baru Telkomsel.',
        NOW(),
        NOW()
    ),
    (
        4,
        'Kartu As adalah produk dari Telkomsel. Kartu ini diluncurkan pada tahun 2004. Kartu As ini memiliki slogan "Pas Buat Semua". Pada tanggal 18 Juni 2021, kartu ini dileburkan bersama simPATI dan LOOP menjadi Telkomsel PraBayar dengan diluncurkannya logo baru Telkomsel.',
        NOW(),
        NOW()
    ),
    (
        5,
        'Telkomsel LOOP sebelumnya bernama simPATI LOOP atau yang dikenali sebagai LOOP adalah kartu GSM prabayar dari Telkomsel yang dirancang untuk kawula muda. Slogan yang dimiliki oleh LOOP adalah "Ini Kita!" Pada tanggal 18 Juni 2021, kartu ini dileburkan bersama simPATI dan Kartu As menjadi Telkomsel PraBayar dengan diluncurkannya logo baru Telkomsel.',
        NOW(),
        NOW()
    ),
    (
        6,
        'Nikmati internetan 24 jam non-stop dengan kuota utama plus bebas nelpon sepuasnya ke sesama IM3 dan Tri selama 5000 menit. Saatnya kamu perluas jangkauan biar bisa ngobrol sama teman kapan pun di mana pun dan makin bebas buat ekspresikan lifestyle kamu lewat konten, sampai nonton YouTube atau main game sepuasnya pakai Freedom Internet dari IM3. Kapan lagi internetan dan nelpon bisa sehemat dan sepuas ini? Yuk, aktifkan segera Freedom Internet di *123#, melalui Aplikasi myIM3, atau IM3 Official WhatsApp!',
        NOW(),
        NOW()
    ),
    (
        7,
        'Sekarang kamu bisa makin puas #TerusTerusan akses aplikasi pakai Freedom U karena ada aplikasi-aplikasi tambahan baru seperti Netflix, Snapchat, Webex, Microsoft Teams, Skype, dan Google Classroom. Plus 24 jam bisa akses aplikasi lebih banyak dengan aplikasi favorit lainnya seperti YouTube, Instagram, TikTok, Facebook, Spotify, Joox, WhatsApp dan Line. Kamu juga bisa pakai Freedom U dengan aplikasi favorit lainnya seperti YouTube, Instagram, TikTok, Facebook, Spotify, Joox, WhatsApp dan Line. Tunggu apa lagi? Aktifkan Freedom U di *123#, myIM3, atau IM3 Ooredoo Official WhatsApp supaya makin puas internetan #TerusTerusan.',
        NOW(),
        NOW()
    ),
    (
        8,
        'Saatnya jalin komunikasi dengan yang terdekat pakai Freedom Combo! Dapatkan kuota besar hingga 60GB dan nelpon sepuasnya ke IM3 Ooredoo dan Tri. Nyaman non-stop telponan dan internetan dengan sahabat, teman dan kerabatmu tetap lancar di mana aja, kapan aja! Chatting, video call jadi bisa sepuasnya, nelpon juga lebih leluasa. Aktifkan di *123#, myIM3, atau IM3 Ooredoo Official WhatsApp.',
        NOW(),
        NOW()
    );
-- 1g. Insert 3 payment methods
INSERT INTO `payment_methods` (`name`, `created_at`, `updated_at`)
VALUES ('Debit/Kredit', NOW(), NOW()),
    ('GoPay', NOW(), NOW()),
    ('OVO', NOW(), NOW());
-- 1h. Insert 5 user pada table users
INSERT INTO `users` (
        `name`,
        `birth_date`,
        `status`,
        `gender`,
        `created_at`,
        `updated_at`
    )
VALUES (
        'Shiyam Lutfiyah Samaha',
        '1959-03-17',
        'KAWIN TERCATAT',
        'PEREMPUAN',
        NOW(),
        NOW()
    ),
    (
        'Hadi Khaldun Koury',
        '1949-08-12',
        'CERAI MATI',
        'LAKI-LAKI',
        NOW(),
        NOW()
    ),
    (
        'Elijah Farnell',
        '1992-04-13',
        'CERAI HIDUP',
        'LAKI-LAKI',
        NOW(),
        NOW()
    ),
    (
        'Brodie Hull',
        '1993-11-7',
        'KAWIN BELUM TERCATAT',
        'LAKI-LAKI',
        NOW(),
        NOW()
    ),
    (
        'Ashley Divine',
        '2000-06-07',
        'BELUM KAWIN',
        'PEREMPUAN',
        NOW(),
        NOW()
    );
-- 1i. Insert 3 transaksi di masing-masing user
INSERT INTO `transactions` (
        `id`,
        `user_id`,
        `total_price`,
        `total_qty`,
        `payment_method_id`,
        `created_at`,
        `updated_at`
    )
VALUES (1, 1, 112000, 3, 1, NOW(), NOW()),
    (2, 1, 117000, 3, 2, NOW(), NOW()),
    (3, 1, 60000, 3, 3, NOW(), NOW()),
    (4, 2, 112000, 3, 1, NOW(), NOW()),
    (5, 2, 117000, 3, 2, NOW(), NOW()),
    (6, 2, 60000, 3, 3, NOW(), NOW()),
    (7, 3, 112000, 3, 1, NOW(), NOW()),
    (8, 3, 117000, 3, 2, NOW(), NOW()),
    (9, 3, 60000, 3, 3, NOW(), NOW()),
    (10, 4, 112000, 3, 1, NOW(), NOW()),
    (11, 4, 117000, 3, 2, NOW(), NOW()),
    (12, 4, 60000, 3, 3, NOW(), NOW()),
    (13, 5, 112000, 3, 1, NOW(), NOW()),
    (14, 5, 117000, 3, 2, NOW(), NOW()),
    (15, 5, 60000, 3, 3, NOW(), NOW());
-- 1j. Insert 3 product di masing-masing transaksi
INSERT INTO `transaction_details` (
        `transaction_id`,
        `product_id`,
        `price`,
        `qty`,
        `created_at`,
        `updated_at`
    )
VALUES (1, 8, 25000, 1, NOW(), NOW()),
    (1, 1, 80000, 1, NOW(), NOW()),
    (1, 6, 7000, 1, NOW(), NOW()),
    (2, 7, 5000, 1, NOW(), NOW()),
    (2, 2, 102000, 1, NOW(), NOW()),
    (2, 5, 10000, 1, NOW(), NOW()),
    (3, 2, 30000, 1, NOW(), NOW()),
    (3, 4, 20000, 1, NOW(), NOW()),
    (3, 5, 10000, 1, NOW(), NOW()),
    (4, 8, 25000, 1, NOW(), NOW()),
    (4, 1, 80000, 1, NOW(), NOW()),
    (4, 6, 7000, 1, NOW(), NOW()),
    (5, 7, 5000, 1, NOW(), NOW()),
    (5, 2, 102000, 1, NOW(), NOW()),
    (5, 5, 10000, 1, NOW(), NOW()),
    (6, 2, 30000, 1, NOW(), NOW()),
    (6, 4, 20000, 1, NOW(), NOW()),
    (6, 5, 10000, 1, NOW(), NOW()),
    (7, 8, 25000, 1, NOW(), NOW()),
    (7, 1, 80000, 1, NOW(), NOW()),
    (7, 6, 7000, 1, NOW(), NOW()),
    (8, 7, 5000, 1, NOW(), NOW()),
    (8, 2, 102000, 1, NOW(), NOW()),
    (8, 5, 10000, 1, NOW(), NOW()),
    (9, 2, 30000, 1, NOW(), NOW()),
    (9, 4, 20000, 1, NOW(), NOW()),
    (9, 5, 10000, 1, NOW(), NOW()),
    (10, 8, 25000, 1, NOW(), NOW()),
    (10, 1, 80000, 1, NOW(), NOW()),
    (10, 6, 7000, 1, NOW(), NOW()),
    (11, 7, 5000, 1, NOW(), NOW()),
    (11, 2, 102000, 1, NOW(), NOW()),
    (11, 5, 10000, 1, NOW(), NOW()),
    (12, 2, 30000, 1, NOW(), NOW()),
    (12, 4, 20000, 1, NOW(), NOW()),
    (12, 5, 10000, 1, NOW(), NOW()),
    (13, 8, 25000, 1, NOW(), NOW()),
    (13, 1, 80000, 1, NOW(), NOW()),
    (13, 6, 7000, 1, NOW(), NOW()),
    (14, 7, 5000, 1, NOW(), NOW()),
    (14, 2, 102000, 1, NOW(), NOW()),
    (14, 5, 10000, 1, NOW(), NOW()),
    (15, 2, 30000, 1, NOW(), NOW()),
    (15, 4, 20000, 1, NOW(), NOW()),
    (15, 5, 10000, 1, NOW(), NOW());
-- SELECT
-- 2a. Tampilkan nama user/pelanggan dengan gender Laki-laki/M
SELECT *
FROM users
WHERE gender = 'LAKI-LAKI';
-- 2b. Tampilkan product dengan id = 3
SELECT *
FROM products
WHERE id = 3;
-- 2c. Tampilkan data pelanggan yang created_at dalam range 7 hari ke belakang dan mempunyai nama mengandung kata 'a'
SELECT *
FROM `users`
where users.name LIKE '%a%'
    AND `created_at` > CURDATE() - INTERVAL 7 DAY;
-- 2d. Hitung jumlah user / pelanggan dengan status gender Perempuan
SELECT *
FROM `users`
WHERE gender = 'PEREMPUAN';
-- 2e. Tampilkan data pelanggan dengan urutan nama sesuai nama abjad
SELECT *
FROM users
ORDER BY `name`;
-- 2f. Tampilkan 5 data pada data product
SELECT *
FROM products
LIMIT 5;
-- UPDATE
-- 3a. Ubah data product id 1 dengan nama 'product dummy'
UPDATE products
SET `name` = 'product dummy'
WHERE `id` = 1;
-- 3b. Update qty = 3 pada transaction detail dengan product id 1
UPDATE transaction_details
SET `qty` = 3
WHERE product_id = 1;
-- DELETE
-- 4a. Delete data pada tabel product dengan id 1
DELETE FROM `products`
WHERE id = 1;
-- 4b. Delete data pada tabel product dengan product type id 1
DELETE FROM `products`
WHERE product_type_id = 1;
-- JOIN, UNION, SUB-QUERY, FUNCTION
-- 1. Gabungkan data transaksi dari user id 1 dan user id 2
select *
from `users`
    join `transactions` on transactions.user_id = users.id
where users.id in (1, 2) -- 2. Tampilkan jumlah harga transaksi user id 1
select sum(total_price)
from transactions
where user_id = 1;
-- 3. Tampilkan total transaksi dengan product type 2
select count(*)
from transactions
where id in (
        select distinct transaction_id
        from transaction_details
        where product_id in (
                select id
                from products
                where product_type_id = 2
            )
    );
-- 4. Tampilkan semua field table product dan field name table product type saling berhubungan
select *
from products
    join product_types on products.product_type_id = product_types.id;
-- 5. Tampilkan semua field table transaction, field name table product dan field name table user
select t.id as transaction_id,
    t.total_price,
    t.total_qty,
    t.payment_method_id,
    t.created_at,
    t.updated_at,
    p.id as product_id,
    p.name as product_name,
    p.product_type_id,
    p.created_at as product_created_at,
    p.updated_at as product_updated_at,
    u.id as user_id,
    u.name as user_name,
    u.birth_date as user_birth_date,
    u.`status` as user_status,
    u.gender as user_gender,
    u.created_at as user_created_at,
    u.updated_at as user_updated_at
from transactions t
    inner join users u on t.user_id = u.id
    inner join transaction_details d on t.id = d.transaction_id
    inner join products p on d.product_id = p.`id`;
-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
DELIMITER $$ CREATE FUNCTION `DeleteTransaction`(transactionId INT) RETURNS tinyint(1) DETERMINISTIC BEGIN
DELETE FROM `transaction_details`
WHERE transaction_id = transactionId;
DELETE FROM `transactions`
WHERE `id` = transactionId;
RETURN true;
END $$ DELIMITER;
-- 7. Buat function setelah data transaksi detail dihapus maka data total qty terupdate berdasarkan qty data transaction id yang dihapus.
DELIMITER $$ CREATE FUNCTION `DeleteTransactionDetail`(transactionDetailId INT) RETURNS tinyint(1) DETERMINISTIC BEGIN
DELETE FROM `transaction_details`
WHERE `id` = transactionDetailId;
UPDATE `transactions`
SET total_qty = (
        SELECT SUM(qty)
        FROM `transaction_details`
        WHERE transaction_id = (
                SELECT transaction_id
                FROM `transaction_details`
                WHERE `id` = transactionDetailId
            )
    )
WHERE `id` = (
        SELECT transaction_id
        FROM `transaction_details`
        WHERE `id` = transactionDetailId
    );
RETURN true;
END $$ DELIMITER;
-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction details dengan sub-query.
select *
from products
where products.id not in (
        select distinct product_id
        from transaction_details
    );