import json
import logging
import os
from pathlib import Path
import shutil
import tempfile
import unittest
from packaging.installer import DeploymentType
from packaging import ApplicationPackage, FunctionAppPackage, CreateApplicationPackageOptions
from packaging._zip_utils import unzip_file


log = logging.getLogger(__name__)

class TestApplicationPackage(unittest.TestCase):
    def setUp(self):
        self.test_dir = tempfile.mkdtemp()
        self.data_dir = Path(os.path.join(os.path.dirname(__file__), 'data', 'app_packaging'))

        self.main_template = self.data_dir / 'templates' / 'main.tf'
        self.create_ui_definition = self.data_dir / 'createUIDefinition.json'
        self.fake_create_ui_definition = self._create_fake_file('fake_create_ui_definition.json')
        self.fake_function_app_package = FunctionAppPackage(self._create_fake_file('fake_function_app_package.zip'))

    def tearDown(self):
        shutil.rmtree(self.test_dir)
        
    def test_main_template(self):
        app_package = ApplicationPackage("main.tf", self.fake_create_ui_definition)
        self.assertEqual(app_package.manifest.main_template, "main.tf")
        self.assertEqual(app_package.manifest.deployment_type, DeploymentType.terraform)

    def test_main_template(self):
        app_package = ApplicationPackage(self.main_template, self.fake_create_ui_definition)
        result = app_package.create(CreateApplicationPackageOptions("v0.0.0"))

        print(result)

        self.assertEqual(app_package.main_template.name, "mainTemplate.json")
        self.assertIsNotNone(app_package.main_template.document)
        self.assertIsNotNone(app_package.main_template.document["variables"]["userData"])

    def test_create(self):
        app_package = ApplicationPackage(self.main_template, self.create_ui_definition)
        options = CreateApplicationPackageOptions(vmi_reference_id = "test-vmi", 
                                                  function_app_package = self.fake_function_app_package)
        result = app_package.create(options)

        self.assertEqual(len(result.validation_results), 0)

        unzip_file(result.file, self.test_dir)
        log.info(os.listdir(self.test_dir))

        main_template_path = Path(self.test_dir).joinpath('mainTemplate.json')
        with open(main_template_path) as main_template_file:
            main_template = json.load(main_template_file)
            self.assertIsNotNone(main_template)
            self.assertIsNotNone(main_template["variables"]["userData"])
            self.assertEqual(len(main_template["variables"]["userData"]["parameters"].keys()), 3)

            self.assertEqual(main_template["variables"][options.vmi_reference_id_variable], options.vmi_reference_id)
            self.assertEqual(main_template["variables"][options.function_app_name_variable], result.function_app_name)
            self.assertEqual(main_template["variables"]["userData"]["installerPackage"]["hash"], result.installer_package.hash)

        # verify the contents of the installer.zip

        shutil.rmtree(result.file.parent)
    
    def _create_fake_file(self, file_name):
        file_path = Path(self.test_dir).joinpath(file_name)
        with open(file_path, 'w') as fp: 
            fp.write('{"parameters": { "outputs": { "secure_variable": "", "location": "", "resource_group_name": "" } }}')
            fp.close()
        return file_path