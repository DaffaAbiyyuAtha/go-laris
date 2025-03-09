WITH inserted_product AS (
    INSERT INTO "product" ("name_product", "price", "discount", "categories_id", "description")
    VALUES
        ('Samsung S24 Ultra', 24000000, 10, 4, 'ini Samsung S24 Ultra'),
        ('PS 5', 10000000, 10, 9, 'ini PS 5'),
        ('Samsung A71', 5700000, 0, 4, 'ini Samsung A71')
    RETURNING id, name_product
)
INSERT INTO "product_images" ("product_id", "image")
SELECT 
    id, 
    unnest(
        CASE name_product
            WHEN 'Samsung S24 Ultra' THEN ARRAY[
                'https://images.samsung.com/is/image/samsung/p6pim/id/2401/gallery/id-galaxy-s24-s928-sm-s928bztqxid-539319760?$624_624_PNG$',
                'https://images.samsung.com/is/image/samsung/assets/id/smartphones/galaxy-s24-ultra/buy/01_S24Ultra-Group-KV_PC_0527_final.jpg',
                'https://images.samsung.com/is/image/samsung/p6pim/id/2401/gallery/id-galaxy-s24-s928-sm-s928bzywxid-539319991?$720_576_JPG$'
            ]
            WHEN 'PS 5' THEN ARRAY[
                'https://gmedia.playstation.com/is/image/SIEPDC/ps5-product-thumbnail-01-en-14sep21?$facebook$',
                'https://datascripmall.id/media/catalog/product/cache/95a5307f46190cd7a50cf0819ebeb220/3/_/3_166.webp',
                'https://asset.kompas.com/crops/NE1EStoz5golGBCK7I5vv91cHYk=/0x0:1440x960/1200x800/data/photo/2023/08/24/64e69aae4768f.jpg'
            ]
            WHEN 'Samsung A71' THEN ARRAY[
                'https://images.samsung.com/is/image/samsung/id-galaxy-a71-a715-sm-a715fzbexid-Blue-211350556?$684_547_PNG$',
                'https://images.samsung.com/is/image/samsung/sa-en-galaxy-a71-a715-sm-a715fzkgksa-front-205720535?$684_547_PNG$',
                'https://images.samsung.com/is/image/samsung/id-galaxy-a71-a715-sm-a715fzbexid-Blue-211350556?$684_547_PNG$'
            ]
        END
    )
FROM inserted_product;
