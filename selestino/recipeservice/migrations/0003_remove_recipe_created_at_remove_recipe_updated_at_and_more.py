
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('recipeservice', '0002_rename_difficulty_level_recipe_difficulty_and_more'),
    ]

    operations = [
        migrations.RemoveField(
            model_name='recipe',
            name='created_at',
        ),
        migrations.RemoveField(
            model_name='recipe',
            name='updated_at',
        ),
        migrations.AlterField(
            model_name='recipe',
            name='cook_time',
            field=models.IntegerField(default=0),
        ),
        migrations.AlterField(
            model_name='recipe',
            name='cuisine_type',
            field=models.CharField(default='Unknown', max_length=100),
        ),
        migrations.AlterField(
            model_name='recipe',
            name='difficulty',
            field=models.CharField(choices=[('Easy', 'Easy'), ('Medium', 'Medium'), ('Hard', 'Hard')], default='Easy', max_length=50),
        ),
        migrations.AlterField(
            model_name='recipe',
            name='ingredients',
            field=models.ManyToManyField(to='recipeservice.ingredient'),
        ),
        migrations.AlterField(
            model_name='recipe',
            name='prep_time',
            field=models.IntegerField(default=0),
        ),
        migrations.AlterField(
            model_name='recipe',
            name='total_time',
            field=models.IntegerField(blank=True, editable=False, null=True),
        ),
    ]
