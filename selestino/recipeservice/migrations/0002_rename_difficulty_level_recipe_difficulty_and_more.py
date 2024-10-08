
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('recipeservice', '0001_initial'),
    ]

    operations = [
        migrations.RenameField(
            model_name='recipe',
            old_name='difficulty_level',
            new_name='difficulty',
        ),
        migrations.AddField(
            model_name='recipe',
            name='cuisine_type',
            field=models.CharField(default='International', max_length=100),
        ),
    ]
