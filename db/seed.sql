-- Seed tables and sample data for PostgreSQL
-- Run with: psql "$DATABASE_URL" -f db/seed.sql

BEGIN;

-- products
CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    brand VARCHAR(100),
    img VARCHAR(255),
    category JSONB,
    description TEXT,
    code VARCHAR(50),
    tag VARCHAR(50),
    links JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- sub_items
CREATE TABLE IF NOT EXISTS sub_items (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    subtitle VARCHAR(100),
    brand VARCHAR(100),
    img VARCHAR(255),
    category JSONB,
    description TEXT,
    code VARCHAR(50),
    shopee_link VARCHAR(255),
    tiktok_link VARCHAR(255),
    lazada_link VARCHAR(255),
    other_link VARCHAR(255),
    display_order SMALLINT DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- highlights
CREATE TABLE IF NOT EXISTS highlights (
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    priority SMALLINT DEFAULT 0,
    end_date TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT unique_product UNIQUE (product_id)
);

-- seed products from JSON samples
WITH data AS (
    SELECT * FROM (VALUES
        ('sillicons-ultra-wide-led-desk-lamp-v2', 'Sillicons - Ultra-Wide LED Desk Lamp V2', '["setup","light"]', 'SILLICONS STUDIO', 'assets/products/sillicons-desk-lamp.png', 'new', 'Ultra-Wide LED Desk Lamp V2 | Wide Lighting Coverage | 4-Color Temperature Modes | 5 Brightness Levels | Touch & Remote Control | Sleek Modern Aesthetic', '{"shopee":"https://s.shopee.co.th/9Kc1XQm17G","tiktok":"","lazada":"https://s.lazada.co.th/s.ZYffBJ","other":""}'),
        ('reptilian-kx78-he', 'Reptilian KX78 HE', '["keyboard"]', 'Saru Space', 'assets/products/saru-kx78he.jpg', 'new', 'Gaming Magnetic Keyboard | Compact 75%| Hot-Swappable | Hall Effect | 8K Polling Rate | Rapid trigger 0.01mm | Tri-mode Connection', '{"shopee":"https://s.shopee.co.th/806LteAKoG","tiktok":"","lazada":"","other":""}'),
        ('itx-build', 'ITX Build', '["PC"]', 'PC Build', 'assets/pc/build.png', 'PC Build', 'Monitor : MSI MAG 274QRFW - 27" IPS 2K 180Hz\nCase: Formd T1 v2.1\nMainboard: Aorus X870i Pro Ice\nCPU: AMD Ryzen 7800X3D\nRAM: 16GB Team DDR5 5600MHz (x2)\nGPU: RTX 4070 super\nSSD: 1TB m.2 KINSTON KC3000\nPSU: Cooler Master SFX 850W\nCPU Cooler: Master Liquid Atmos ii 240 (White)', '{}'),

        ('razer-phantom-white', 'Razer - Phantom White', '["collection"]', 'razer', 'assets/products/collection-razer-phantom-white.jpg', '', '', '{}'),
        ('elgato-neo', 'Elgato Neo', '["collection"]', 'elgato', 'assets/products/collection-elgato-neo.jpg', '', '', '{}'),
        ('saru-space-reptilian-kx78-he', 'Saru Space - Reptilian KX78 HE', '["keyboard"]', 'Saru Space', 'assets/products/saru-kx78he.jpg', '', 'Gaming Magnetic Keyboard | Compact 75%| Hot-Swappable | Hall Effect | 8K Polling Rate | Rapid trigger 0.01mm | Tri-mode Connection', '{"shopee":"https://s.shopee.co.th/806LteAKoG","tiktok":"","lazada":"","other":""}'),
        ('womier-era75', 'Womier ERA75', '["keyboard"]', 'womier', 'assets/products/womier-era75.jpg', '', '[Womier] 75% mechanical keyboard kit | hot-swappable | gasket mount', '{"shopee":"","tiktok":"","lazada":"","other":"https://womierkeyboard.com/products/womier-era75?sca_ref=10149033.MaOc1663RR"}'),
        ('ergonoz-kb01', 'Ergonoz KB01', '["keyboard"]', 'ergonoz', 'assets/products/ergonoz-kb01.png', '', 'Ergonoz Keyboard Kb01 Is a 2-System Wireless Keyboard That Supports Both 2.4Ghz and Bluetooth.', '{"shopee":"https://s.shopee.co.th/7KpG6XQ4KR","tiktok":"https://vt.tiktok.com/ZSHE18R4VYk5S-03YDL/","lazada":"https://s.lazada.co.th/s.ZbjyUb","other":""}'),
        ('gconic-a75-pro', 'Gconic A75 Pro', '["keyboard"]', 'gconic', 'assets/products/gconic-a75-pro.png', '', '', '{"shopee":"https://s.shopee.co.th/1LXvt3QGIg","tiktok":"","lazada":"https://s.lazada.co.th/s.ZbjBgw","other":""}'),
        ('gconic-a98-pro', 'Gconic A98 Pro', '["keyboard"]', 'gconic', 'assets/products/gconic-a98-pro.png', '', '', '{"shopee":"https://s.shopee.co.th/5ffE7IeBeb","tiktok":"","lazada":"","other":""}'),
        ('ajazz-nk68', 'Ajazz NK68', '["keyboard"]', 'ajazz', 'assets/products/ajazz-nk68.png', '', '', '{"shopee":"https://s.shopee.co.th/7fTBKcTkXU","tiktok":"https://vt.tiktok.com/ZSHctV9VqXm5B-novoI/","lazada":"https://s.lazada.co.th/s.ZbjzbZ","other":""}'),
        ('ajazz-nk61', 'Ajazz NK61', '["keyboard"]', 'ajazz', 'assets/products/ajazz-nk61.png', '', '', '{"shopee":"https://s.shopee.co.th/1qSYiHKIQ4","tiktok":"https://vt.tiktok.com/ZSHE12etDnxgw-nHSRD/","lazada":"","other":""}'),
        ('ajazz-ak980-wired', 'Ajazz AK980 Wired', '["keyboard"]', 'ajazz', 'assets/products/ajazz-ak980-wired.png', '', '', '{"shopee":"https://s.shopee.co.th/30dqbSPIBw","tiktok":"","lazada":"","other":""}'),
        ('qwertykeys-hex80', 'Hex80', '["keyboard"]', 'qwertykeys x atk', 'assets/products/qwertykeys-hex80.png', '', '', '{"shopee":"https://s.shopee.co.th/6fYF75MXns","tiktok":"","lazada":"","other":""}'),
        ('qwertykeys-qk80-mk2', 'QK80 MK ii', '["keyboard"]', 'qwertykeys', 'assets/products/qwertykeys-qk80-mk2.png', '', '', '{"shopee":"","tiktok":"","lazada":"","other":""}'),
        ('royalkludge-s98', 'RK S98', '["keyboard"]', 'royalkludge', 'assets/products/royalkludge-s98.png', '', '', '{"shopee":"https://s.shopee.co.th/4Ap6W5GJKJ","tiktok":"","lazada":"","other":""}'),
        ('elecfox-inky75', 'INKY 75', '["keyboard"]', 'elecfox', 'assets/products/elecfox-inky75.png', '', '', '{"shopee":"https://s.shopee.co.th/2B2oQqgDNq","tiktok":"","lazada":"","other":""}'),
        ('royalkludge-m3', 'RK M3', '["mouse"]', 'royalkludge', 'assets/products/royalkludge-m3.png', '', '', '{"shopee":"https://s.shopee.co.th/LeDhg77JE","tiktok":"","lazada":"","other":""}'),
        ('mambasnake-m5-ultra', 'M5 Ultra', '["mouse"]', 'mambasnake', 'assets/products/mambasnake-m5-ultra.png', '', '', '{"shopee":"https://s.shopee.co.th/LeDhg77JE","tiktok":"","lazada":"","other":""}'),
        ('logitech-mx-master-3', 'Logitech MX Master 3', '["mouse"]', 'logitech', 'assets/products/logitech-mx-master-3.png', '', '', '{"shopee":"https://s.shopee.co.th/6VHHImzk3m","tiktok":"","lazada":"","other":""}'),
        ('ajazz-aj159-apex', 'Ajazz AJ159 APEX', '["mouse"]', 'ajazz', 'assets/products/ajazz-aj159-apex.png', '', '', '{"shopee":"https://s.shopee.co.th/2B8Iz5E60i","tiktok":"https://vt.tiktok.com/ZSHEjnn4cdeju-O67ur/","lazada":"","other":""}'),
        ('orbitkey-deskmat', 'Orbitkey Desk Mat', '["mousepad","deskmat"]', 'orbitkey', 'assets/products/orbitkey-deskmat.png', '', '', '{"shopee":"https://s.shopee.co.th/4L2H0pKMzJ","tiktok":"","lazada":"","other":""}'),
        ('divoom-times-frame', 'Divoom Times Frame', '["gadgets"]', 'divoom', 'assets/products/divoom-times-frame.png', '', '', '{"shopee":"","tiktok":"https://vt.tiktok.com/ZSHG4nowoLnBV-w4n5I/","lazada":"","other":""}'),
        ('divoom-tiivoo2', 'Divoom Tiivoo2', '["gadgets","speakers"]', 'divoom', 'assets/products/divoom-tiivoo2.png', '', '', '{"shopee":"","tiktok":"https://vt.tiktok.com/ZSD3KGWaF/","lazada":"","other":""}'),
        ('divoom-ditoo', 'Divoom Ditoo', '["gadgets","speakers"]', 'divoom', 'assets/products/divoom-ditoo.png', '', '', '{"shopee":"https://s.shopee.co.th/6KjPXgXLZB","tiktok":"","lazada":"","other":""}'),
        ('vinko-gan-k2002', 'Vinko GaN K2002', '["gadgets","adapter"]', 'vinko', 'assets/products/vinko-gan-k2002.png', '', 'Power Adapter 65W', '{"shopee":"","tiktok":"https://vt.tiktok.com/ZSHsxMn8nBh1J-Uoxdq/","lazada":"","other":""}'),
        ('vinko-v200', 'Vinko V200', '["gadgets","powerbank"]', 'vinko', 'assets/products/vinko-v200.png', '', 'Power Bank 130W', '{"shopee":"","tiktok":"https://vt.tiktok.com/ZSH3mKPRuD33E-BH1Wv/","lazada":"","other":""}'),
        ('ugreen-uno-6-in-1', 'Ugreen Uno 6-in-1', '["gadgets"]', 'ugreen', 'assets/products/ugreen-uno-6-in-1.png', '', '', '{"shopee":"https://s.shopee.co.th/3VSzXRg26M","tiktok":"","lazada":"","other":""}'),
        ('jasoz-usb-hub', 'Jasoz USB Hub', '["gadgets"]', 'jasoz', 'assets/products/jasoz-usb-hub.png', '', '', '{"shopee":"https://s.shopee.co.th/7ANj1BVODm","tiktok":"","lazada":"","other":""}'),
        ('ergonoz-pyroz', 'Ergonoz - Pyroz', '["setup"]', 'ergonoz', 'assets/products/ergonoz-pyroz.jpg', '', 'Ergonoz Monitor Arm | 17"-32" | 2-9 kg | VESA 75x75mm / 100x100mm | Machanical Spring | Cable Management', '{"shopee":"https://s.shopee.co.th/9pY054fbTY","tiktok":"https://vt.tiktok.com/ZSHoJYDy5VL2T-rjnVP/","lazada":"https://s.lazada.co.th/s.ZbE5nc","other":""}'),
        ('diy-desk-stand', 'Desk Stand', '["setup"]', '', 'assets/products/diy-desk-stand.png', '', '', '{"shopee":"https://s.shopee.co.th/1VrZRomPCr","tiktok":"","lazada":"","other":""}'),
        ('picture-frame', 'Picture Frame', '["setup"]', '', 'assets/products/picture-frame.png', '', '', '{"shopee":"https://s.shopee.co.th/AUjPlfRfuu","tiktok":"","lazada":"","other":""}'),
        ('stand-generic', 'Stand', '["setup"]', '', 'assets/products/stand.png', '', '', '{"shopee":"https://s.shopee.co.th/3LEUC6cTtC","tiktok":"","lazada":"","other":""}'),
        ('fifine-bm88', 'Fifine BM88', '["setup"]', '', 'assets/products/fifine-bm88.png', '', '', '{"shopee":"https://s.shopee.co.th/9KWU5mTDSH","tiktok":"","lazada":"","other":""}'),
        ('xiaomi-monitor-lightbar', 'Xiaomi Monitor LightBar', '["setup"]', '', 'assets/products/xiaomi-monitor-lightbar.png', '', '', '{"shopee":"https://s.shopee.co.th/9ABiBze4cD","tiktok":"","lazada":"","other":""}'),
        ('table-bewell', 'Table Bewell', '["setup"]', '', 'assets/products/table-bewell.png', '', '', '{"shopee":"https://s.shopee.co.th/6pkjOmRaXB","tiktok":"","lazada":"","other":""}'),
        ('bewell-frozen', 'Bewell FROZEN', '["setup"]', '', 'assets/products/bewell-frozen.png', '', '', '{"shopee":"https://s.shopee.co.th/4fgEp0qW0H","tiktok":"","lazada":"","other":""}'),
        ('board-decorations', 'Board Decorations', '["setup"]', '', 'assets/products/board-decorations.png', '', '', '{"shopee":"https://s.shopee.co.th/8UolJeS6PD","tiktok":"","lazada":"","other":""}'),
        ('picture-frame-ps2-controller', 'Picture frame PS2 controller', '["setup"]', '', 'assets/products/picture-frame-ps2-controller.png', '', '', '{"shopee":"https://s.shopee.co.th/9UpN1Ce6AG","tiktok":"","lazada":"","other":""}'),
        ('skadis-pegboard', 'Skadis Pegboard', '["setup"]', '', 'assets/products/skadis-pegboard.png', '', '', '{"shopee":"https://s.shopee.co.th/7AJNj2OV2A","tiktok":"","lazada":"","other":""}'),
        ('fifine-k688', 'Fifine K688', '["microphone"]', '', 'assets/products/fifine-k688.png', '', '', '{"shopee":"https://link.fifinedesign.com/K688-kenzchiro","tiktok":"","lazada":"","other":""}'),
        ('fifine-am8', 'FIFINE AM8', '["microphone"]', '', 'assets/products/fifine-am8.png', '', '', '{"shopee":"https://s.shopee.co.th/VthHOE3W6","tiktok":"","lazada":"","other":""}'),
        ('maono-dm40', 'Maono dm40', '["microphone"]', '', 'assets/products/maono-dm40.png', '', '', '{"shopee":"https://s.shopee.co.th/4q3eEpaygV","tiktok":"","lazada":"","other":""}'),
        ('comica-ejoy-uni-pro', 'Comica Ejoy Uni Pro', '["microphone"]', '', 'assets/products/comica-ejoy-uni-pro.png', '', '', '{"shopee":"https://s.shopee.co.th/AUj74HtRdO","tiktok":"","lazada":"","other":""}'),
        ('boya-magic-mic-4-in-1', 'Boya Magic Mic 4 in 1', '["microphone"]', '', 'assets/products/boya-magic-mic-4-in-1.png', '', '', '{"shopee":"","tiktok":"https://vt.tiktok.com/ZSD37YrM3/","lazada":"","other":""}'),
        ('maono-t1-mini', 'Maono T1 Mini', '["microphone"]', '', 'assets/products/maono-t1-mini.png', '', '', '{"shopee":"","tiktok":"https://vt.tiktok.com/ZSD3WH6J5/","lazada":"","other":""}'),
        ('xiaomi-sound-party', 'Xiaomi Sound party', '["headphone"]', '', 'assets/products/xiaomi-sound-party.png', '', '', '{"shopee":"","tiktok":"https://vt.tiktok.com/ZSDcMJcCB/","lazada":"","other":""}'),
        ('kanto-ora', 'Kanto Ora', '["headphone"]', '', 'assets/products/kanto-ora.png', '', '', '{"shopee":"","tiktok":"https://vt.tiktok.com/ZSHWEwtKY8JpS-CC7qa/","lazada":"","other":""}'),
        ('corsair-hs80-wireless', 'CORSAIR HEADSET HS80 WIRELESS', '["headphone"]', '', 'assets/products/corsair-headset.png', '', '', '{"shopee":"https://s.shopee.co.th/2fza46vuy5","tiktok":"","lazada":"","other":""}'),
        ('soundcore-aeroclip', 'Soundcore AeroClip', '["headphone"]', '', 'assets/products/soundcore-aeroclip.png', '', '', '{"shopee":"","tiktok":"https://vt.tiktok.com/ZSDT1V1P4/","lazada":"","other":""}'),
        ('torras-ostand-q3-air', 'Torras - Ostand Q3 Air', '["phone accessories"]', 'torras', 'assets/products/torras-ostand-q3-air.png', '', 'iPhone Case', '{"shopee":"https://s.shopee.co.th/4LBpb9gqgd","tiktok":"https://vt.tiktok.com/ZSHE1nwr55Wu7-T4JGv/","lazada":"","other":""}'),
        ('maktar-qubii-duo', 'Qubii duo', '["phone accessories"]', 'maktar', 'assets/products/maktar-qubii-duo.png', '', '', '{"shopee":"","tiktok":"https://vt.tiktok.com/ZSHcBScmJqxfd-jF6uH/","lazada":"","other":""}')
    ) AS t(code, title, category, brand, img, tag, description, links)
)
INSERT INTO products (title, category, brand, img, tag, description, code, links)
SELECT title,
       category::jsonb,
       NULLIF(brand, ''),
       img,
       NULLIF(tag, ''),
       description,
       code,
       links::jsonb
FROM data
WHERE NOT EXISTS (SELECT 1 FROM products p WHERE p.code = data.code);

-- seed sub_items (from group_items / ITX build)
WITH parent AS (
    SELECT code, id FROM products WHERE code IN ('itx-build','razer-phantom-white','elgato-neo')
), data AS (
    SELECT * FROM (VALUES
        -- ITX Build group
        ('itx-build', 'Formd T1 (mini ITX)', 'Case', '["pc","case"]', 'formd', 'assets/pc/case.png', '9.95L SFF case | A4 paper-sized footprint | AIO support | CNC machined aluminum | Premium anodizing', 'formd-t1-mini-itx', '{"shopee":"https://s.shopee.co.th/2AytqlipMP","tiktok":"","lazada":"","other":"https://formdt1.com/?srsltid=AfmBOoqru1hRu31Qvmv-Bl1yMeR3Ly9d5QGz91FZ37Ao1V1_Xk1P1T_v"}', 1),
        ('itx-build', 'Aorus - X870I AORUS PRO ICE', 'Mainboard', '["pc","mainboard"]', 'aorus', 'assets/pc/mainboard.png', 'AMD X870I chipset | 8+2+1 phase VRM | DDR5 8400+ MT/s | PCIe 5.0 x16 | 2x M.2 (PCIe 5.0/4.0) | USB4 40Gbps | WiFi 7 | 2.5GbE LAN', 'aorus-x870i-aorus-pro-ice', '{"shopee":"https://s.shopee.co.th/3VXMhtMcGp","tiktok":"","lazada":"","other":"https://www.gigabyte.com/th/Motherboard/X870I-AORUS-PRO-ICE-rev-10"}', 2),
        ('itx-build', 'AMD - Ryzen 7 7800X3D', 'CPU', '["pc","cpu"]', 'amd', 'assets/pc/cpu.png', '8 cores 16 threads | Up to 5.0 GHz boost | 96MB 3D V-Cache | Zen 4 architecture | 120W TDP | Dominant gaming processor', 'amd-ryzen-7-7800x3d', '{"shopee":"https://s.shopee.co.th/4VPq3hwj9U","tiktok":"","lazada":"","other":""}', 3),
        ('itx-build', 'Cooler Master - Master Liquid Atmos 240 ii - LCD', 'CPU Cooler', '["pc","cpu cooler"]', 'cooler master', 'assets/pc/cpu-cooler.png', '240mm AIO liquid cooler | LCD screen display | Dual 120mm fans | Optimized for high-end CPUs | Silent operation', 'cooler-master-master-liquid-atmos-240-ii-lcd', '{"shopee":"https://s.shopee.co.th/AKUKaclio3","tiktok":"","lazada":"","other":""}', 4),
        ('itx-build', 'MSI MAG 274QRFW - 27" IPS 2K 180Hz', 'Monitor', '["pc","monitor"]', 'msi', 'assets/pc/monitor.png', '27" IPS panel | 2560x1440 (2K) | 180Hz refresh rate | Rapid IPS 1ms response | HDR support | Gaming features', 'msi-mag-274qrfw', '{"shopee":"https://s.shopee.co.th/1Vohpmfyfj","tiktok":"","lazada":"","other":""}', 5),
        -- Razer Phantom White collection
        ('razer-phantom-white', 'Barracuda - Phantom White', NULL, '["phantom white","headset"]', 'razer', 'assets/products/razer-barracuda.png', '', 'razer-barracuda-phantom-white', '{"shopee":"https://s.shopee.co.th/30fsbEEnEC","tiktok":"","lazada":"","other":""}', 1),
        ('razer-phantom-white', 'Basilisk V3 Pro - Phantom White', NULL, '["phantom white","mouse"]', 'razer', 'assets/products/razer-basilisk-v3-pro.png', '', 'razer-basilisk-v3-pro-phantom-white', '{"shopee":"https://s.shopee.co.th/50Qwyvp7GJ","tiktok":"","lazada":"","other":""}', 2),
        ('razer-phantom-white', 'Black Widow V4 - Phantom White', NULL, '["phantom white","keyboard"]', 'razer', 'assets/products/razer-black-widow-v4.png', '', 'razer-black-widow-v4-phantom-white', '{"shopee":"https://s.shopee.co.th/804YYWF7wO","tiktok":"","lazada":"","other":""}', 3),
        ('razer-phantom-white', 'Firefly V2 - Phantom White', NULL, '["phantom white","mousepad"]', 'razer', 'assets/products/razer-firefly-v2.png', '', 'razer-firefly-v2-phantom-white', '{"shopee":"https://s.shopee.co.th/VyXcnSnNm","tiktok":"","lazada":"","other":""}', 4),
        -- Elgato Neo collection
        ('elgato-neo', 'Elgato - Wave Neo', NULL, '["neo","microphone"]', 'elgato', 'assets/products/elgato-wave-neo.png', '', 'elgato-wave-neo', '{"shopee":"https://s.shopee.co.th/1qVPklXv0u","tiktok":"https://vt.tiktok.com/ZSHEdExtBHKTk-FilJI/","lazada":"https://s.lazada.co.th/s.Zb96Sc","other":""}', 1),
        ('elgato-neo', 'Elgato - Stream Deck Neo', NULL, '["neo","gadgets"]', 'elgato', 'assets/products/elgato-stream-deck-neo.png', '', 'elgato-stream-deck-neo', '{"shopee":"https://s.shopee.co.th/8KitUj73Dt","tiktok":"https://vt.tiktok.com/ZSHEdExtBHKTk-FilJI/","lazada":"https://s.lazada.co.th/s.Zb96Sc","other":""}', 2),
        ('elgato-neo', 'Elgato - Capture Card Neo', NULL, '["neo","gadgets"]', 'elgato', 'assets/products/elgato-capture-card-neo.png', '', 'elgato-capture-card-neo', '{"shopee":"https://s.shopee.co.th/9pXhHcxuZ2","tiktok":"https://vt.tiktok.com/ZSHEdExtBHKTk-FilJI/","lazada":"https://s.lazada.co.th/s.Zb96Sc","other":""}', 3),
        ('elgato-neo', 'Elgato - Facecam Neo', NULL, '["neo","webcam"]', 'elgato', 'assets/products/elgato-facecam-neo.png', '', 'elgato-facecam-neo', '{"shopee":"https://s.shopee.co.th/8KitV1zRdc","tiktok":"https://vt.tiktok.com/ZSHEdExtBHKTk-FilJI/","lazada":"https://s.lazada.co.th/s.Zb96Sc","other":""}', 4)
    ) AS t(parent_code, title, subtitle, category, brand, img, description, code, links, display_order)
)
INSERT INTO sub_items (product_id, title, subtitle, brand, img, category, description, code, shopee_link, tiktok_link, lazada_link, other_link, display_order)
SELECT p.id,
       d.title,
       d.subtitle,
       NULLIF(d.brand, ''),
       d.img,
       d.category::jsonb,
       d.description,
       d.code,
       d.links::jsonb->>'shopee',
       d.links::jsonb->>'tiktok',
       d.links::jsonb->>'lazada',
       d.links::jsonb->>'other',
       d.display_order
FROM data d
JOIN parent p ON p.code = d.parent_code
WHERE NOT EXISTS (SELECT 1 FROM sub_items s WHERE s.code = d.code);


-- seed highlights
WITH data AS (
    SELECT * FROM (VALUES
        ('sillicons-ultra-wide-led-desk-lamp-v2', 1),
        ('reptilian-kx78-he', 2),
        ('itx-build', 3)
    ) AS t(code, priority)
)
INSERT INTO highlights (product_id, priority, end_date)
SELECT p.id, d.priority, now() + interval '30 days'
FROM data d
JOIN products p ON p.code = d.code
WHERE NOT EXISTS (SELECT 1 FROM highlights h WHERE h.product_id = p.id);

COMMIT;
